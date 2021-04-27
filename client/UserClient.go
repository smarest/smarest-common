package client

import (
	"net/http"

	"github.com/smarest/smarest-common/client/resource"
)

type UserClient struct {
	EntryPoint string
	*POSClient
}

func NewUserClient(host string, timeout int) *UserClient {
	return &UserClient{"/v1/user", NewPOSClient(host, timeout)}
}

func (c *UserClient) GetUser(requestResource *resource.RequestResource) (*http.Response, error) {
	url := c.Host + c.EntryPoint

	return c.DoPostRequest(url, requestResource)
}
