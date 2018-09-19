package util

import (
	"io"
	"net/http"
)

func SendRequest(url string, contentType string, body io.Reader) (*http.Response, error) {
	resp, err := http.Post(url, contentType, body)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
