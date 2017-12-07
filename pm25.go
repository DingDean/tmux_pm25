package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

type Air struct {
	Pm2_5   int    `json:"pm2_5"`
	Area    string `json:"area"`
	Quality string `json:"quality"`
}

type Geo struct {
	City string `json:"city"`
}

type Aircache struct {
	City      string
	Content   string
	Timestamp int64
}

var netClient = &http.Client{
	Timeout: time.Second * 10,
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func get_config() string {
	viper.SetConfigName(".tmux_25_config")
	viper.AddConfigPath("$HOME/.tmux_25_config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	check(err)
	return viper.GetString("apiKey")
}

func get_city_name() string {
	res, err := netClient.Get("http://freegeoip.net/json/")
	check(err)
	defer res.Body.Close()
	buf, err := ioutil.ReadAll(res.Body)
	check(err)
	var body Geo
	err = json.Unmarshal(buf, &body)
	check(err)
	return strings.ToLower(body.City)
}

func get_pm25() (Air, string, error) {
	city := get_city_name()
	apiKey := get_config()
	url := fmt.Sprintf("http://www.pm25.in/api/querys/pm2_5.json?city=%s&stations=no&token=%s", city, apiKey)
	res, err := netClient.Get(url)
	check(err)
	defer res.Body.Close()
	if res.StatusCode != 200 {
		panic(res.Status)
	}
	buf, err := ioutil.ReadAll(res.Body)
	check(err)
	var body []Air
	err = json.Unmarshal(buf, &body)
	check(err)
	return body[0], city, nil
}

func isExpired(then int64) bool {
	now := time.Now().Unix()
	diff := now - then
	return diff > 3600 // 缓存只有一个小时的时效
}

func check_cache(cacheFilepath string) string {
	dat, err := ioutil.ReadFile(cacheFilepath)
	if err != nil {
		return ""
	}
	var jsdat Aircache
	if err = json.Unmarshal(dat, &jsdat); err != nil {
		panic(err)
	}
	if isExpired(jsdat.Timestamp) {
		return ""
	}
	return jsdat.Content
}

func main() {
	cacheFilepath, err := filepath.Abs("./.tmux_25_cache")
	check(err)

	data := check_cache(cacheFilepath)
	if data == "" {
		raw, city, _ := get_pm25()
		data = fmt.Sprintf("%s %d %s\n", raw.Area, raw.Pm2_5, raw.Quality)
		cache := Aircache{
			City:      city,
			Content:   data,
			Timestamp: time.Now().Unix(),
		}
		jscache, err := json.Marshal(cache)
		check(err)
		defer ioutil.WriteFile(cacheFilepath, jscache, 0644)
	}
	fmt.Println(data)
}
