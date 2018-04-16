package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"
)

type Aqi struct {
	Pm2_5   string
	Area    string
	Quality string
	Error   string
	Time    int64
}

func (r Aqi) Echo() {
	if r.Error != "" {
		fmt.Printf("出错啦")
		return
	}
	fmt.Printf("%s %s %s", r.Area, r.Pm2_5, r.Quality)
}

type AqiService interface {
	Query(city string) Aqi
}

func main() {
	cacheFilepath, err := filepath.Abs("./.tmux_25_cache")
	if err != nil {
		panic(err.Error())
	}
	data := getCache(cacheFilepath)
	if !cacheIsExpired(data.Time) {
		data.Echo()
		return
	}
	conf := getConf(".tmux_25_config")

	var api AqiService
	source := conf.Source
	Appcode := conf.ApiKey
	Req := http.Client{Timeout: time.Second * 10}
	if source == "aliyun" {
		api = Aliyun{Req: Req, Appcode: Appcode}
	} else if source == "pm25.in" {
		api = Pm25In{Req: Req, Appcode: Appcode}
	} else {
		panic("未知的API源")
	}
	city := conf.City
	aqiData := api.Query(city)
	saveCache(aqiData, cacheFilepath)
	aqiData.Echo()
}
