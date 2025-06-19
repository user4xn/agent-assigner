package helper

import (
	"io"
	"log"
	"net/http"
)

// helper for closing request client
func ClientClose(client *http.Response) {
	if err := client.Body.Close(); err != nil {
		log.Print(err)
	}
}

// helper for get request
func GetRequest(client *http.Client, url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, body)
	if err != nil {
		return &http.Response{}, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	response, err := client.Do(req)
	if err != nil {
		return &http.Response{}, err
	}

	return response, nil
}

// helper for post request
func PostRequest(client *http.Client, url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return &http.Response{}, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	response, err := client.Do(req)
	if err != nil {
		return &http.Response{}, err
	}

	return response, nil
}
