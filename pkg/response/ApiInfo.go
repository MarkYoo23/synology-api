package response

// ApiInfo represents the response from the SYNO.API.Info endpoint
// {"data":{"SYNO.API.Auth":{"maxVersion":7,"minVersion":1,"path":"entry.cgi"}},"success":true}
type ApiInfo struct {
	Data ApiInfoData `json:"data"`
}

type ApiInfoData struct {
	Auth    ApiInfoAuth `json:"SYNO.API.Auth"`
	Success bool        `json:"success"`
}

type ApiInfoAuth struct {
	MaxVersion int    `json:"maxVersion"`
	MinVersion int    `json:"minVersion"`
	Path       string `json:"path"`
}
