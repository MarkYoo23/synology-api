package client

import (
	"encoding/json"
	"errors"
	"net/http"
	"synology-api/pkg/request"
	"synology-api/pkg/response"
)

func (c *Client) GetApiInfo() (*response.ApiInfo, error) {
	b := request.UriBuilder{
		CgiPath: request.ApiCgiQuery,
		ApiName: request.ApiNameInfo,
		Version: 1,
		Method:  request.ApiMethodQuery,
		Params: map[request.ApiParam]string{
			request.ApiParamQuery: "all",
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
