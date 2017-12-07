package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Air struct {
	Pm2_5             int    `json:"pm2_5"`
	Primary_pollutant string `json:"primary_pollutant"`
	Quality           string `json:"quality"`
}

type Aircache struct {
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

func get_pm25() (Air, error) {
	city := "hangzhou"
	url := fmt.Sprintf("http://www.pm25.in/api/querys/pm2_5.json?city=%s&stations=no&token=5j1znBVAsnSf5xQyNQyq", city)
	res, err := netClient.Get(url)
	check(err)
	defer res.Body.Close()
	buf, err := ioutil.ReadAll(res.Body)
	check(err)
	var body []Air
	err = json.Unmarshal(buf, &body)
	check(err)
	return body[0], nil
}

func isExpired(then int64) bool {
	now := time.Now().Unix()
	diff := now - then
	return diff > 3600000 // 缓存只有一个小时的时效
}

func check_cache() string {
	dat, err := ioutil.ReadFile("./cache")
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
	data := check_cache()
	if data == "" {
		raw, _ := get_pm25()
		data = fmt.Sprintf("%d %s\n", raw.Pm2_5, raw.Quality)
		cache := Aircache{
			Content:   data,
			Timestamp: time.Now().Unix(),
		}
		jscache, err := json.Marshal(cache)
		check(err)
		defer ioutil.WriteFile("./cache", jscache, 0644)
	}
	fmt.Println(data)
}
