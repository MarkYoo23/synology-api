package synology_api

import (
	"errors"
	"github.com/MarkYoo23/synology-api/request"
	"io"
	"net/http"
)

func (c *Client) OpenFile(path string) (io.ReadCloser, error) {
	b := c.makeFileBuilder(path, "open")

	statusCode, rc, err := c.getIO(b)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, errors.New("failed to open file")
	}

	return rc, nil
}

func (c *Client) DownloadFile(path string) (io.ReadCloser, error) {
	b := c.makeFileBuilder(path, "download")

	statusCode, rc, err := c.getIO(b)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, errors.New("failed to open file")
	}

	return rc, nil
}

func (c *Client) makeFileBuilder(path string, mode string) request.UriBuilder {
	b := request.UriBuilder{
		CgiPath: request.ApiCgiEntry,
		ApiName: request.ApiNameFileDownload,
		Version: 2,
		Method:  request.ApiMethodFileDownload,
		Sid:     c.sessionId,
		Params: map[request.ApiParam]string{
			request.ApiParamDownloadPath: path,
			request.ApiParamDownloadMode: mode,
		},
	}

	return b
}
