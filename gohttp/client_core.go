package gohttp

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

const (
	defaultMaxIdleConnections = 5
	defaultResponseTimeout    = 5 * time.Second
	defaultConnectionTimeout  = 1 * time.Second
)

func (c *httpClient) getRequestHeaders(requestHeaders http.Header) http.Header {
	result := make(http.Header)
	// add default headers to the request
	for header, value := range c.builder.headers {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	// add custom headers to the request
	for header, value := range requestHeaders {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	return result
}

func (c *httpClient) getRequestBody(contentType string, body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}

	switch strings.ToLower(contentType) {
	case "application/json":
		return json.Marshal(body)

	case "application/xml":
		return xml.Marshal(body)

	default:
		return json.Marshal(body)
	}
}

func (c *httpClient) getMaxIdleConnections() int {
	if c.builder.maxIdleConnections > 0 {
		return c.builder.maxIdleConnections
	}

	return defaultMaxIdleConnections
}

func (c *httpClient) getResponseTimeout() time.Duration {
	if c.builder.responseTimeout > 0 {
		return c.builder.responseTimeout
	}

	if c.builder.disableTimeout {
		return 0
	}

	return defaultResponseTimeout
}

func (c *httpClient) getConnectionTimeout() time.Duration {
	if c.builder.connectionTimeout > 0 {
		return c.builder.connectionTimeout
	}

	if c.builder.disableTimeout {
		return 0
	}

	return defaultResponseTimeout
}

func (c *httpClient) getHttpClient() *http.Client {
	c.clientOnce.Do(func() {
		fmt.Println("= = = = = = = = = = = = = = = = = = = = = = = ")
		fmt.Println("CREATING A NEW HTTP CLIENT!")
		fmt.Println("= = = = = = = = = = = = = = = = = = = = = = = ")
		c.client = &http.Client{
			Timeout: c.getConnectionTimeout() + c.getResponseTimeout(),
			Transport: &http.Transport{
				MaxIdleConnsPerHost: c.getMaxIdleConnections(),
				// the maximum amount of time that it's going to wait for a response when we send the request.
				ResponseHeaderTimeout: c.getResponseTimeout(),
				DialContext: (&net.Dialer{
					// the maximum amount of time that it's going to wait for a given connection
					Timeout: c.getConnectionTimeout(),
				}).DialContext,
			},
		}
	})
	return c.client
}

func (c *httpClient) do(httpMethod string, url string, headers http.Header, body interface{}) (*Response, error) {
	fullHeaders := c.getRequestHeaders(headers)

	requestBody, err := c.getRequestBody(fullHeaders.Get("Content-Type"), body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(httpMethod, url, bytes.NewBuffer(requestBody))

	if err != nil {
		return nil, errors.New("unable to create a new request")
	}

	request.Header = fullHeaders

	client := c.getHttpClient()

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	bodyResponse, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	finalResponse := Response{
		status:     response.Status,
		statusCode: response.StatusCode,
		headers:    response.Header,
		body:       bodyResponse,
	}
	return &finalResponse, nil
}
