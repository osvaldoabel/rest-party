package gohttp

import (
	"net/http"
	"sync"
)

type Client interface {
	SetHeaders(headers http.Header)
	Get(url string, headers http.Header) (*Response, error)
	Post(url string, headers http.Header, body interface{}) (*Response, error)
	Put(url string, headers http.Header, body interface{}) (*Response, error)
	Patch(url string, headers http.Header, body interface{}) (*Response, error)
	Delete(url string, headers http.Header) (*Response, error)
}

type httpClient struct {
	builder    *clientBuilder
	client     *http.Client
	clientOnce sync.Once
}

func (c *httpClient) SetHeaders(headers http.Header) {
	c.builder.headers = headers
}

func (c *httpClient) DisableTimeouts(disable bool) {
	c.builder.disableTimeout = disable
}

func (c *httpClient) Get(url string, headers http.Header) (*Response, error) {
	return c.do(http.MethodGet, url, headers, nil)
}

func (c *httpClient) Post(url string, headers http.Header, body interface{}) (*Response, error) {
	return c.do(http.MethodPost, url, headers, nil)
}

func (c *httpClient) Put(url string, headers http.Header, body interface{}) (*Response, error) {
	return c.do(http.MethodPut, url, headers, nil)

}

func (c *httpClient) Patch(url string, headers http.Header, body interface{}) (*Response, error) {
	return c.do(http.MethodPatch, url, headers, nil)
}

func (c *httpClient) Delete(url string, headers http.Header) (*Response, error) {
	return c.do(http.MethodDelete, url, headers, nil)
}
