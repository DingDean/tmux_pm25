package aqi

import (
	"encoding/json"
	"errors"
	"fmt"
	pm25 "github.com/DingDean/tmux_pm25"
	"net/http"
)

type AliRes struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
	Result struct {
		City string `json:"city"`
		Aqi  struct {
			Pm2_5   string `json:"pm2_5"`
			Quality string `json:"quality"`
		} `json:"aqi"`
	} `json:"result"`
}

type Aliyun struct {
	Req     http.Client
	Appcode string
}

func (x Aliyun) Query(city string) pm25.Aqi {
	// 创建请求
	url := fmt.Sprintf("http://jisutqybmf.market.alicloudapi.com/weather/query?city=%s", city)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return makeErrorAqi(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("APPCODE %s", x.Appcode))
	// 执行请求
	res, err := x.Req.Do(req)
	if err != nil {
		return makeErrorAqi(err)
	}
	defer res.Body.Close()
	body, err := parseBody(res)
	if err != nil {
		return makeErrorAqi(err)
	}
	var aqidata AliRes
	err = json.Unmarshal(body, &aqidata)
	if err != nil {
		return makeErrorAqi(err)
	}
	// 检查返回值
	if aqidata.Status != "0" {
		return makeErrorAqi(errors.New(aqidata.Msg))
	}
	return pm25.Aqi{
		Pm2_5:   aqidata.Result.Aqi.Pm2_5,
		Area:    aqidata.Result.City,
		Quality: aqidata.Result.Aqi.Quality,
		Error:   "",
	}
}
