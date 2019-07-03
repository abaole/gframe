package ghttp

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/abaole/gframe/logger"
)

func GET(url string) (map[string]string, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	response, err := client.Get(url)
	if err != nil {
		logger.Errorf(" GET http.NewRequest %+v ", err)
		return nil, errors.New("request url error")
	}
	//程序在使用完回复后必须关闭回复的主体。
	defer response.Body.Close()

	body, err1 := ioutil.ReadAll(response.Body)
	if err1 != nil {
		logger.Errorf(" GET 请求原信息  %v", err1)
		return nil, errors.New("request url error")
	}

	var result map[string]string
	errJson := json.Unmarshal(body, &result)
	if errJson != nil {
		logger.Errorf(" json.Unmarshal %+v ", errJson)
		return nil, errors.New("request url error")
	}
	return result, nil
}

func PostJSON(url string, param []byte) ([]byte, error) {
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(param))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json;charset=utf-8;")

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Errorf("ioutil.ReadAll %+v ", err)
		return nil, err
	}

	return content, nil
}

func Get(c context.Context, uri string) (resp *http.Response, err error) {
	done := make(chan *http.Response, 1)
	errChan := make(chan error, 1)

	if reps, err := http.NewRequest("GET", uri, nil); err != nil {
		return nil, err
	} else {
		go func() {
			if resq, err := http.DefaultClient.Do(reps); err != nil {
				errChan <- err
			} else {
				done <- resq
			}
		}()

		select {
		case err := <-errChan:
			return nil, err
		case resp := <-done:
			return resp, nil
		}
	}
}

func Post(c context.Context, uri string, param map[string]string) (resp *http.Response, err error) {
	jsonByte, _ := json.Marshal(param)
	done := make(chan *http.Response, 1)
	errChan := make(chan error, 1)

	if reps, err := http.NewRequest("POST", uri, bytes.NewReader(jsonByte)); err != nil {
		return nil, err
	} else {
		go func() {
			if resq, err := http.DefaultClient.Do(reps); err != nil {
				errChan <- err
			} else {
				done <- resq
			}
		}()

		select {
		case err := <-errChan:
			return nil, err
		case resp := <-done:
			return resp, nil
		}
	}
}
