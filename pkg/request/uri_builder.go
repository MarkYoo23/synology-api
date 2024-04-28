package request

import (
	"fmt"
	"net/url"
)

// UriBuilder webapi/<CGI_PATH>?api=<API_NAME>&version=<VERSION>&method=<METHOD>[&<PARAMS>][&_sid=<SID>]
type UriBuilder struct {
	CgiPath ApiCgi
	ApiName ApiName
	Version int
	Method  ApiMethod
	Params  map[ApiParam]string
	Sid     string
}

func (u UriBuilder) Build() string {
	query := url.Values{}

	query.Set("api", string(u.ApiName))
	query.Set("version", fmt.Sprintf("%d", u.Version))
	query.Set("method", string(u.Method))

	for k, v := range u.Params {
		query.Set(string(k), v)
	}

	if u.Sid != "" {
		query.Set("_sid", u.Sid)
	}

	return fmt.Sprintf("/webapi/%s?%s", u.CgiPath, query.Encode())
}
