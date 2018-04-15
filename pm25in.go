package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Pm25In struct {
	Req     http.Client
	Appcode string
}

func (x Pm25In) Query(city string) Aqi {
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
	var body []Aqi
	err = json.Unmarshal(buf, &body)
	if err != nil {
		return makeErrorAqi(err)
	}
	return body[0]
}
