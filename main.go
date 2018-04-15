package main

import (
	"filepath"
	"fmt"
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

type Conf struct {
	City   string
	ApiKey string
	Source string
}

type AqiService interface {
	Query(city string) Aqi
}

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
