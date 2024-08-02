package helper

import (
	"fmt"
	"github.com/spf13/viper"
)

func LoadEnv() *viper.Viper {

	v := viper.New()

	v.AddConfigPath("config")
	v.SetConfigName("auth-config")
	v.SetConfigType("env")
	v.AutomaticEnv()
	err := v.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf(err.Error()))
	}

	return v
}
