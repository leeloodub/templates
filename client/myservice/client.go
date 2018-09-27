package myservice

import (
	"encoding/json"
	"net/http"
	"net/url"

	sm "../messages"
)

const (
	get  = "GET"
	post = "POST"
)

// API endpoints
const (
	fooUrl = "/v1/foo"
)

type Client struct {
	baseURL    *url.URL
	httpClient *http.Client
}

func NewClient(url *url.URL) *Client {
	return &Client{
		baseURL:    url,
		httpClient: http.DefaultClient,
	}
}

type Header struct {
	key   string
	value string
}

func (c *Client) DoSomething(payload sm.Request) (sm.Response, error) {
	rel := &url.URL{Path: fooUrl}
	u := c.baseURL.ResolveReference(rel)

	req, err := http.NewRequest("METHOD", u.String(), nil)
	if err != nil {
		return sm.Response{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return sm.Response{}, err
	}
	defer resp.Body.Close()

	var decoded sm.Response
	err = json.NewDecoder(resp.Body).Decode(&decoded)
	return decoded, err
}
