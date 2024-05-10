package synology_api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MarkYoo23/synology-api/request"
	"github.com/MarkYoo23/synology-api/response"
	"net/http"
)

func (c *Client) Login(id, pw string) error {
	b := request.UriBuilder{
		CgiPath: request.ApiCgiAuthV1to3,
		ApiName: request.ApiNameAuth,
		Version: 3,
		Method:  request.ApiMethodLogin,
		Params: map[request.ApiParam]string{
			request.ApiParamAccount: id,
			request.ApiParamPasswd:  pw,
			request.ApiParamSession: c.serverName,
			request.ApiParamFormat:  "sid",
		},
	}

	statusCode, bytes, err := c.get(b)
	if err != nil {
		return err
	}

	if statusCode != http.StatusOK {
		return errors.New("failed to login")
	}

	var r response.LoginV3
	err = json.Unmarshal(bytes, &r)
	if err != nil {
		return err
	}

	c.deviceId = r.Data.Did
	c.sessionId = r.Data.Sid
	return nil
}

func (c *Client) IsLogin() bool {
	return c.sessionId != "" && c.deviceId != ""
}

func (c *Client) Logout() error {
	b := request.UriBuilder{
		CgiPath: request.ApiCgiAuthV1to3,
		ApiName: request.ApiNameAuth,
		Version: 1,
		Method:  request.ApiMethodLogout,
		Params: map[request.ApiParam]string{
			request.ApiParamSession: c.serverName,
		},
		Sid: c.sessionId,
	}

	statusCode, bytes, err := c.get(b)
	if err != nil {
		return err
	}

	if statusCode != http.StatusOK {
		return errors.New("failed to logout")
	}

	fmt.Printf("Logged out : %s\n", string(bytes))

	c.deviceId = ""
	c.sessionId = ""
	return nil
}
