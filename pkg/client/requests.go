package client

import (
	"io/ioutil"
	"net/http"
)

func get(url string) ([]byte, *http.Response, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, nil, err
	}

	return body, resp, nil
}
