package config

import (
	"os"

	"go.uber.org/zap"
)

var Logger *zap.Logger

func InitLogger() {
	var err error

	os.Mkdir("/var/log/mikronet/", os.ModePerm)

	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"/var/log/mikronet/app.log"}
	Logger, err = config.Build()

	if err != nil {
		panic("failed to initialize logger: " + err.Error())
	}
}
