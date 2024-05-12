package synology_api

import (
	"github.com/MarkYoo23/synology-api/request"
	"io"
	"net/http"
	"time"
)

type Client struct {
	serverName string
	host       string

	// if the client is authenticated
	deviceId  string
	sessionId string
}

func NewClient(serverName, url string) *Client {
	return &Client{
		serverName: serverName,
		host:       url,
	}
}

func (c *Client) get(builder request.UriBuilder) (statusCode int, body []byte, err error) {
	clt := c.newHttpClient(time.Second * 10)

	url := c.host + builder.Build()
	r, err := clt.Get(url)
	if err != nil {
		return 404, nil, err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(r.Body)

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		return r.StatusCode, nil, err
	}

	return r.StatusCode, bytes, nil
}

func (c *Client) getIO(builder request.UriBuilder) (statusCode int, closer io.ReadCloser, err error) {
	clt := c.newHttpClient(time.Second * 10)

	url := c.host + builder.Build()
	r, err := clt.Get(url)
	if err != nil {
		return 404, nil, err
	}

	return r.StatusCode, r.Body, nil
}

func (c *Client) newHttpClient(timeout time.Duration) *http.Client {
	return &http.Client{
		Timeout: timeout,
	}
}
