package helper

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func LoadEnv() *viper.Viper {

	v := viper.New()

	path, _ := os.Getwd()
	dir := filepath.Dir(path)
	v.SetConfigName("auth-config")
	v.SetConfigType("env")
	v.AddConfigPath(dir + "/config")

	err := v.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf(err.Error()))
	}

	return v
}
