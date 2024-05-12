package response

type ApiInfoData map[string]struct {
	MaxVersion    int    `json:"maxVersion"`
	MinVersion    int    `json:"minVersion"`
	Path          string `json:"path"`
	RequestFormat string `json:"requestFormat"`
}

type ApiInfo struct {
	ApiInfoData `json:"data"`
	Success     bool `json:"success"`
}
