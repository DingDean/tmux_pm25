package main

import (
	"errors"
	"io/ioutil"
	"net/http"
)

func makeErrorAqi(err error) pm25.Aqi {
	return pm25.Aqi{Error: err.Error()}
}

func parseBody(res *http.Response) ([]byte, error) {
	var body []byte
	if res.StatusCode != 200 {
		return body, errors.New(res.Status)
	}
	// 读取Body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return body, err
	}
	return body, nil
}
