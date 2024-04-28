package response

type LoginV3 struct {
	Data struct {
		Did string `json:"did"`
		Sid string `json:"sid"`
	} `json:"data"`
	Success bool `json:"success"`
}
