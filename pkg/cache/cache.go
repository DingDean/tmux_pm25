package cache

import (
	"encoding/json"
	"fmt"
	pm25 "github.com/DingDean/tmux_pm25"
	"io/ioutil"
	"time"
)

func Get(cacheFilepath string) pm25.Aqi {
	dat, err := ioutil.ReadFile(cacheFilepath)
	if err != nil {
		var zeros pm25.Aqi
		return zeros
	}
	var jsdat pm25.Aqi
	if err = json.Unmarshal(dat, &jsdat); err != nil {
		panic(fmt.Sprintf("无法获取缓存, %s", err.Error()))
	}
	return jsdat
}

func IsExpired(then int64) bool {
	now := time.Now().Unix()
	diff := now - then
	return diff > 3600 // 缓存只有一个小时的时效
}

func Save(cache pm25.Aqi, cacheFilepath string) {
	cache.Time = time.Now().Unix()
	data, err := json.Marshal(cache)
	if err != nil {
		data = []byte(fmt.Sprintf("{Time: %d, Error: %s}", time.Now().Unix(), err.Error()))
	}
	ioutil.WriteFile(cacheFilepath, data, 0644)
}
