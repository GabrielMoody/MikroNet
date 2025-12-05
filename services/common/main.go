package common

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type AMQP struct {
	URL string

	conn    *amqp.Connection
	channel *amqp.Channel
	ackCh   chan amqp.Confirmation
	notify  chan *amqp.Error
}

func New(url string) *AMQP {
	return &AMQP{URL: url}
}

func (a *AMQP) Connect(ctx context.Context) error {
	var err error
	backoff := time.Second
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		a.conn, err = amqp.Dial(a.URL)
		if err == nil {
			a.channel, err = a.conn.Channel()
			if err == nil {
				a.notify = a.conn.NotifyClose(make(chan *amqp.Error))
				// enable confirm mode for publisher reliability
				_ = a.channel.Confirm(false)
				a.ackCh = a.channel.NotifyPublish(make(chan amqp.Confirmation, 1))

				return nil
			}
			_ = a.conn.Close()
		}
		log.Printf("amqp connect failed: %v; retrying in %s", err, backoff)
		time.Sleep(backoff)
		if backoff < 30*time.Second {
			backoff *= 2
		}
	}
}

func (a *AMQP) Close() {
	if a.channel != nil {
		_ = a.channel.Close()
	}
	if a.conn != nil {
		_ = a.conn.Close()
	}
}

// DeclareExchange ensures exchange exists (durable)
func (a *AMQP) DeclareExchange(name, kind string) error {
	if a.channel == nil {
		return errors.New("channel nil")
	}
	return a.channel.ExchangeDeclare(name, kind, true, false, false, false, nil)
}

func (a *AMQP) PublishPersistent(exchange, routingKey string, body []byte) error {
	if a.channel == nil {
		return errors.New("channel nil")
	}

	err := a.channel.PublishWithContext(context.Background(), exchange, routingKey, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         body,
		Timestamp:    time.Now(),
	})
	if err != nil {
		return err
	}

	select {
	case conf := <-a.ackCh:
		if !conf.Ack {
			return fmt.Errorf("message nacked")
		}
	case <-time.After(5 * time.Second):
		return fmt.Errorf("publish confirm timeout")
	}
	return nil
}

func (a *AMQP) Consume(queue, exchange, routingKey string) (<-chan amqp.Delivery, error) {
	if a.channel == nil {
		return nil, errors.New("channel nil")
	}
	_, err := a.channel.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		return nil, err
	}
	if exchange != "" {
		if err := a.channel.QueueBind(queue, routingKey, exchange, false, nil); err != nil {
			return nil, err
		}
	}
	msgs, err := a.channel.Consume(queue, "", false, false, false, false, nil)
	if err != nil {
		return nil, err
	}
	return msgs, nil
}
