package gohttp

import (
	"net/http"
	"time"
)

type clientBuilder struct {
	maxIdleConnections int
	connectionTimeout  time.Duration
	responseTimeout    time.Duration
	disableTimeout     bool
	headers            http.Header
}

type ClientBuilder interface {
	SetMaxIdleConnections(i int) ClientBuilder
	SetConnectionTimeout(timeout time.Duration) ClientBuilder
	SetResponseTimeout(timeout time.Duration) ClientBuilder
	DisableTimeouts(disable bool) ClientBuilder
	Build() Client
}

func NewBuilder() ClientBuilder {
	builder := &clientBuilder{}
	return builder
}

func (c *clientBuilder) Build() Client {
	client := httpClient{
		builder: c,
	}

	return &client
}

func (c *clientBuilder) SetMaxIdleConnections(i int) ClientBuilder {
	c.maxIdleConnections = i
	return c
}

func (c *clientBuilder) SetConnectionTimeout(timeout time.Duration) ClientBuilder {
	c.connectionTimeout = timeout
	return c
}

func (c *clientBuilder) SetResponseTimeout(timeout time.Duration) ClientBuilder {
	c.responseTimeout = timeout
	return c
}

func (c *clientBuilder) DisableTimeouts(disable bool) ClientBuilder {
	c.disableTimeout = disable
	return c
}
