package mq

import (
	"encoding/json"
	"github.com/GabrielMoody/mikroNet/driver/internal/dto"
	"github.com/streadway/amqp"
	"time"
)

type MQ struct {
	Conn *amqp.Connection
	Ch   *amqp.Channel
}

func (m *MQ) MQInit() (err error) {
	m.Conn, err = amqp.Dial("amqp://mikronet:123@localhost:5672/")

	if err != nil {
		return err
	}

	defer m.Conn.Close()

	m.Ch, err = m.Conn.Channel()

	if err != nil {
		return err
	}
	defer m.Ch.Close()

	return nil
}

func (m *MQ) PublishMessage(location dto.LocationReq, qname string) (err error) {
	q, err := m.Ch.QueueDeclare(
		qname,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	jsonLoc, _ := json.Marshal(location)

	for {
		err = m.Ch.Publish(
			"",
			q.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        jsonLoc,
			},
		)

		if err != nil {
			return err
		}

		time.Sleep(5 * time.Second)
	}
}
