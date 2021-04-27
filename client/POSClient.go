package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type POSClient struct {
	Host   string
	Client *http.Client
}

func NewPOSClient(host string, timeout int) *POSClient {
	return &POSClient{
		Host: host,
		Client: &http.Client{
			Timeout: time.Millisecond * time.Duration(timeout)},
	}
}

func (c *POSClient) DoGetRequest(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("can not create http request. [%w]", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Send request
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client has error. [%w]", err)
	}

	return resp, nil
}

func (c *POSClient) DoPostRequest(url string, requestResource interface{}) (*http.Response, error) {
	return c.exchange(url, "POST", &requestResource)
}

func (c *POSClient) DoPutRequest(url string, requestResource interface{}) (*http.Response, error) {
	return c.exchange(url, "PUT", &requestResource)
}

func (c *POSClient) exchange(url string, method string, requestResource interface{}) (*http.Response, error) {
	var requestBody *bytes.Buffer = nil
	if requestResource != nil {
		body, err := json.Marshal(requestResource)
		if err != nil {
			return nil, fmt.Errorf("can not create request. [%w]", err)
		}
		requestBody = bytes.NewBuffer(body)
	}

	req, err := http.NewRequest(method, url, requestBody)
	if err != nil {
		return nil, fmt.Errorf("can not create request. [%w]", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Send request
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client hass error. [%w]", err)
	}

	return resp, nil
}
