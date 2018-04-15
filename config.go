package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func Get(name string) pm25.Conf {
	viper.SetConfigName(name)
	viper.AddConfigPath("$HOME")
	viper.AddConfigPath("$GOPATH/src/github.com/DingDean/tmux_pm25/")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic("无法获取配置")
	}
	keys := []string{"city", "apiKey", "source"}
	for _, key := range keys {
		if viper.GetString(key) == "" {
			panic(fmt.Sprintf("缺少配置%s", key))
		}
	}
	return pm25.Conf{
		City:   viper.GetString("city"),
		ApiKey: viper.GetString("apiKey"),
		Source: viper.GetString("source"),
	}
}
