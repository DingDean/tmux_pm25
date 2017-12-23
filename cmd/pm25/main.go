package main

import (
	pm25 "github.com/DingDean/tmux_pm25"
	"github.com/DingDean/tmux_pm25/pkg/aqi"
	"github.com/DingDean/tmux_pm25/pkg/cache"
	"github.com/DingDean/tmux_pm25/pkg/config"
	"net/http"
	"path/filepath"
	"time"
)

func main() {
	cacheFilepath, err := filepath.Abs("./.tmux_25_cache")
	if err != nil {
		panic(err.Error())
	}
	data := cache.Get(cacheFilepath)
	if !cache.IsExpired(data.Time) {
		data.Echo()
		return
	}
	conf := config.Get(".tmux_25_config")

	var api pm25.AqiService
	source := conf.Source
	Appcode := conf.ApiKey
	Req := http.Client{Timeout: time.Second * 10}
	if source == "aliyun" {
		api = aqi.Aliyun{Req: Req, Appcode: Appcode}
	} else if source == "pm25.in" {
		api = aqi.Pm25In{Req: Req, Appcode: Appcode}
	} else {
		panic("未知的API源")
	}
	city := conf.City
	aqiData := api.Query(city)
	cache.Save(aqiData, cacheFilepath)
	aqiData.Echo()
}
