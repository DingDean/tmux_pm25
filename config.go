package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type Conf struct {
	City   string
	ApiKey string
	Source string
}

func getConf(name string) Conf {
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
	return Conf{
		City:   viper.GetString("city"),
		ApiKey: viper.GetString("apiKey"),
		Source: viper.GetString("source"),
	}
}
