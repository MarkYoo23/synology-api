package client

import (
	"encoding/json"
	"errors"
	"net/http"
	"synology-api/pkg/request"
	"synology-api/pkg/response"
)

func (c *Client) GetShareFolder() (*response.ShareFolder, error) {
	b := request.UriBuilder{
		CgiPath: request.ApiCgiEntry,
		ApiName: request.ApiNameFileList,
		Version: 2,
		Method:  request.ApiMethodShareFolderList,
		Sid:     c.sessionId,
	}

	statusCode, bytes, err := c.get(b)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, errors.New("failed to retrieve share folder")
	}

	var r response.ShareFolder
	err = json.Unmarshal(bytes, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (c *Client) GetFolder(path string) (*response.Folder, error) {
	b := request.UriBuilder{
		CgiPath: request.ApiCgiEntry,
		ApiName: request.ApiNameFileList,
		Version: 2,
		Method:  request.ApiMethodFolderList,
		Params: map[request.ApiParam]string{
			request.ApiParamFolderPath: path,
		},
		Sid: c.sessionId,
	}

	statusCode, bytes, err := c.get(b)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, errors.New("failed to retrieve share folder")
	}

	var r response.Folder
	err = json.Unmarshal(bytes, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (c *Client) GetFolderWithAdditional(path string) (*response.FolderWithAdditional, error) {
	b := request.UriBuilder{
		CgiPath: request.ApiCgiEntry,
		ApiName: request.ApiNameFileList,
		Version: 2,
		Method:  request.ApiMethodFolderList,
		Params: map[request.ApiParam]string{
			request.ApiParamFolderPath:       path,
			request.ApiParamFolderAdditional: "[\"real_path\",\"size\",\"owner\",\"time,perm\",\"type\"]",
		},
		Sid: c.sessionId,
	}

	statusCode, bytes, err := c.get(b)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, errors.New("failed to retrieve share folder")
	}

	var r response.FolderWithAdditional
	err = json.Unmarshal(bytes, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
