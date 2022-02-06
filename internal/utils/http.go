package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

func GetHttp(url string, payload []byte) ([]byte, int, error) {
	client := http.Client{
		Timeout: 50 * time.Second,
	}

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(payload))
	if err != nil {
		return []byte{}, 500, &InternalServerError{ErrMessage: err.Error()}
	}

	res, err := client.Do(req)
	if err != nil {
		return []byte{}, 500, &InternalServerError{ErrMessage: err.Error()}
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, 500, &InternalServerError{ErrMessage: err.Error()}
	}
	return body, res.StatusCode, nil
}
