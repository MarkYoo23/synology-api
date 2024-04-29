package response

type Folder struct {
	Data struct {
		Files []struct {
			Isdir bool   `json:"isdir"`
			Name  string `json:"name"`
			Path  string `json:"path"`
		} `json:"files"`
		Offset int `json:"offset"`
		Total  int `json:"total"`
	} `json:"data"`
	Success bool `json:"success"`
}

type FolderWithAdditional struct {
	Data struct {
		Files []struct {
			Additional struct {
				Owner struct {
					Gid   int    `json:"gid"`
					Group string `json:"group"`
					Uid   int    `json:"uid"`
					User  string `json:"user"`
				} `json:"owner"`
				RealPath string `json:"real_path"`
				Size     int    `json:"size"`
				Type     string `json:"type"`
			} `json:"additional"`
			Isdir bool   `json:"isdir"`
			Name  string `json:"name"`
			Path  string `json:"path"`
		} `json:"files"`
		Offset int `json:"offset"`
		Total  int `json:"total"`
	} `json:"data"`
	Success bool `json:"success"`
}
