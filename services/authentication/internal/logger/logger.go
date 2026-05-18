package logger

import (
	"os"

	"github.com/rs/zerolog"
)

func New() zerolog.Logger {
	file, err := os.OpenFile("/var/log/app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		panic(err)
	}

	return zerolog.New(file).
		With().
		Timestamp().
		Str("service", "authentication").
		Logger()
}
