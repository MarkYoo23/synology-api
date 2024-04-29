package response

type ShareFolder struct {
	Data struct {
		Offset int `json:"offset"`
		Shares []struct {
			Isdir bool   `json:"isdir"`
			Name  string `json:"name"`
			Path  string `json:"path"`
		} `json:"shares"`
		Total int `json:"total"`
	} `json:"data"`
	Success bool `json:"success"`
}
