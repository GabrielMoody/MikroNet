package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/GabrielMoody/MikroNet/services/order/config/rabbitmq"
	"github.com/GabrielMoody/MikroNet/services/order/internal/handler"
	"github.com/GabrielMoody/MikroNet/services/order/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db := model.DatabaseInit()
	rdb := model.RedisConnect()
	aqmp_cons := rabbitmq.Init("amqp://admin:admin123@rabbitmq:5672/")
	amqp_pub := rabbitmq.Init("amqp://admin:admin123@rabbitmq:5672/")

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowHeaders: "Authorization, Content-Type",
		AllowOrigins: "*",
	}))
	app.Use(logger.New(logger.Config{
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Asia/Singapore",
	}))

	api := app.Group("/")

	orderEvents := handler.OrderHandler(api, db, rdb, aqmp_cons, amqp_pub)

	if err := orderEvents.Listen(ctx); err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := app.Listen(":8060"); err != nil {
			log.Fatal(err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig

	log.Info("Shutting down...")

	cancel()
	_ = app.Shutdown()
}
