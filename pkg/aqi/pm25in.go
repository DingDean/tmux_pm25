package aqi

import (
	"encoding/json"
	"fmt"
	pm25 "github.com/DingDean/tmux_pm25"
	"net/http"
)

type Pm25In struct {
	Req     http.Client
	Appcode string
}

func (x Pm25In) Query(city string) pm25.Aqi {
	url := fmt.Sprintf("http://www.pm25.in/api/querys/pm2_5.json?city=%s&stations=no&token=%s", city, x.Appcode)
	res, err := x.Req.Get(url)
	if err != nil {
		return makeErrorAqi(err)
	}
	defer res.Body.Close()
	buf, err := parseBody(res)
	if err != nil {
		return makeErrorAqi(err)
	}
	var body []pm25.Aqi
	err = json.Unmarshal(buf, &body)
	if err != nil {
		return makeErrorAqi(err)
	}
	return body[0]
}
