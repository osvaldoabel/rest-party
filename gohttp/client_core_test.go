package gohttp

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T) {
	// Initialization
	client := httpClient{}
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "cool-http-client")
	client.Headers = commonHeaders

	// Execution
	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "ABC-123")

	// Validation
	finalheaders := client.getRequestHeaders(requestHeaders)
	if len(finalheaders) != 3 {
		t.Error("we expect 3 headers")
	}

	if finalheaders.Get("X-Request-Id") != "ABC-123" {
		t.Error("invalid request id received")
	}

	if finalheaders.Get("Content-Type") != "application/json" {
		t.Error("invalid content type received")
	}

	if finalheaders.Get("User-Agent") != "cool-http-client" {
		t.Error("invalid User Agent received")
	}
}

func TestGetRequestBody(t *testing.T) {
	client := httpClient{}

	//initialization
	t.Run("noBodyNilResponse", func(t *testing.T) {
		body, err := client.getRequestBody("", nil)
		if err != nil {
			t.Error("no error expected when passing nil body")
		}

		if body != nil {
			t.Error("no body expected when passing nil body")
		}
	})

	t.Run("BodyWithJson", func(t *testing.T) {
		requestBody := []string{"one", "two"}
		body, err := client.getRequestBody("application/json", requestBody)
		fmt.Println(err)
		fmt.Println(string(body))

		if err != nil {
			t.Error("no error expected when marshaling slice as json")
		}

		if string(body) != `["one","two"]` {
			t.Error("invalid body obtained")
		}
	})

	t.Run("BodyWithXml", func(t *testing.T) {

	})

	t.Run("BodyWithJson", func(t *testing.T) {

	})

	t.Run("BodyWithJsonAsDefault", func(t *testing.T) {

	})
}
