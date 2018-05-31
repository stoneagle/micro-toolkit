package common

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

type Rres struct {
	Result int
	Data   interface{}
	Msg    string
}

func DoHttpPost(url string, params []byte, timeout time.Duration) (body []byte, err error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(params))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: time.Duration(timeout),
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	return
}
