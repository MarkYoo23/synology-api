package response

type ApiInfo struct {
	Data struct {
		SYNOAPIAuth struct {
			MaxVersion int    `json:"maxVersion"`
			MinVersion int    `json:"minVersion"`
			Path       string `json:"path"`
		} `json:"SYNO.API.Auth"`
	} `json:"data"`
	Success bool `json:"success"`
}
