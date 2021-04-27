package client

import (
	"fmt"
	"net/http"
)

type LoginClient struct {
	UserEntryPoint       string
	RestaurantEntryPoint string
	*POSClient
}

func NewLoginClient(host string, timeout int) *LoginClient {
	return &LoginClient{"/v1/login/user", "/v1/login/restaurant", NewPOSClient(host, timeout)}
}

func (c *LoginClient) GetUserByCookie(cookie string) (*http.Response, error) {
	url := fmt.Sprintf("%s%s/%s", c.Host, c.UserEntryPoint, cookie)

	return c.DoGetRequest(url)
}

func (c *LoginClient) PostRestaurant(requestResource interface{}) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", c.Host, c.RestaurantEntryPoint)

	return c.exchange(url, "POST", &requestResource)
}
