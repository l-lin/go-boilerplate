package http

import (
	"fmt"
	"github.com/henvic/httpretty"
	"io/ioutil"
	"net/http"
)

// Client is a wrapper to factorize http calls
type Client struct {
	*http.Client
}

func NewClient(debug bool) *Client {
	client := &http.Client{}
	if debug {
		logger := &httpretty.Logger{
			Time:           true,
			TLS:            true,
			RequestHeader:  true,
			RequestBody:    true,
			ResponseHeader: true,
			ResponseBody:   true,
			Colors:         true,
			Formatters:     []httpretty.Formatter{&httpretty.JSONFormatter{}},
		}
		client.Transport = logger.RoundTripper(http.DefaultTransport)
	}
	return &Client{Client: client}
}

// DoOnlyOk invokes http.Client.Do and return http.Response only if response status code is OK
func (c Client) DoOnlyOk(request *http.Request) (*http.Response, error) {
	response, err := c.Client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		defer response.Body.Close()
		var body []byte
		body, err = ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("http response status was %d and http response body was %s", response.StatusCode, string(body))
	}
	return response, nil
}
