package tmux_pm25

import (
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
