package client

import (
	"fmt"
	"net/http"
)

type LoginClient struct {
	UserEntryPoint string
	*BaseClient
}

func NewLoginClient(host string, timeout int) *LoginClient {
	return &LoginClient{"/v1/login/user", NewBaseClient(host, timeout)}
}

func (c *LoginClient) GetUserByCookie(cookie string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v1/cookie/%s/user", c.Host, cookie)

	return c.DoGetRequest(url)
}
