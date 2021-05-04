package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type BaseClient struct {
	Host   string
	Client *http.Client
}

func NewBaseClient(host string, timeout int) *BaseClient {
	return &BaseClient{
		Host: host,
		Client: &http.Client{
			Timeout: time.Millisecond * time.Duration(timeout)},
	}
}

func (c *BaseClient) DoGetRequest(url string) (*http.Response, error) {
	log.Printf("url[%s] method[get]", url)
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

func (c *BaseClient) DoPostRequest(url string, requestResource interface{}) (*http.Response, error) {
	return c.exchange(url, "POST", &requestResource)
}

func (c *BaseClient) DoPutRequest(url string, requestResource interface{}) (*http.Response, error) {
	return c.exchange(url, "PUT", &requestResource)
}

func (c *BaseClient) exchange(url string, method string, requestResource interface{}) (*http.Response, error) {
	log.Printf("url[%s] method[%s] body[%+v]", url, method, requestResource)
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
