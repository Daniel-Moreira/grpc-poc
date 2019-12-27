package http

import (
	"io"
	"net/http"
)

func Request(method string, url string) (*io.ReadCloser, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return &resp.Body, nil
}
