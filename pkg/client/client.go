package client

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"synology-api/pkg/request"
	"synology-api/pkg/response"
	"time"
)

type Client struct {
	host string
}

func NewClient(url string) *Client {
	return &Client{
		host: url,
	}
}

func (c *Client) GetApiInfo() (*response.ApiInfo, error) {
	b := request.UriBuilder{
		CgiPath: request.ApiCgiQuery,
		ApiName: request.ApiNameInfo,
		Version: 1,
		Method:  request.ApiMethodQuery,
		Params: map[request.ApiParam]string{
			request.ApiParamQuery: "SYNO.API.Auth,SYNO.FileStation",
		},
	}

	status, bytes, err := c.get(b)
	if err != nil {
		return nil, err
	}

	if status != http.StatusOK {
		return nil, errors.New("failed to retrieve API")
	}

	var r response.ApiInfo

	err = json.Unmarshal(bytes, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (c *Client) get(builder request.UriBuilder) (statusCode int, body []byte, err error) {
	clt := c.newHttpClient(time.Second * 3)

	r, err := clt.Get(c.host + builder.Build())
	if err != nil {
		return r.StatusCode, nil, err
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

func (c *Client) newHttpClient(timeout time.Duration) *http.Client {
	return &http.Client{
		Timeout: timeout,
	}
}
