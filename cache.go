package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

func getCache(cacheFilepath string) Aqi {
	dat, err := ioutil.ReadFile(cacheFilepath)
	if err != nil {
		var zeros Aqi
		return zeros
	}
	var jsdat Aqi
	if err = json.Unmarshal(dat, &jsdat); err != nil {
		panic(fmt.Sprintf("无法获取缓存, %s", err.Error()))
	}
	return jsdat
}

func cacheIsExpired(then int64) bool {
	now := time.Now().Unix()
	diff := now - then
	return diff > 3600 // 缓存只有一个小时的时效
}

func saveCache(cache Aqi, cacheFilepath string) {
	cache.Time = time.Now().Unix()
	data, err := json.Marshal(cache)
	if err != nil {
		data = []byte(fmt.Sprintf("{Time: %d, Error: %s}", time.Now().Unix(), err.Error()))
	}
	ioutil.WriteFile(cacheFilepath, data, 0644)
}
