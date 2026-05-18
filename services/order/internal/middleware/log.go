package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

func LoggerMiddleware(base zerolog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		requestID := uuid.New().String()

		logger := base.With().
			Str("request_id", requestID).
			Str("method", c.Method()).
			Str("path", c.Path()).
			Str("ip", c.IP()).
			Logger()

		c.Locals("logger", &logger)

		err := c.Next()

		status := c.Response().StatusCode()

		event := logger.With().
			Int("status", status).
			Dur("latency_ms", time.Since(start)).
			Logger()

		switch {
		case status >= 500:
			event.Error().
				Err(err).
				Msg("http_request")

		case status >= 400:
			event.Warn().
				Msg("http_request")

		default:
			event.Info().
				Msg("http_request")
		}

		return err
	}
}
