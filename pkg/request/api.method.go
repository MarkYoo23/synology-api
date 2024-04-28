package request

type ApiMethod string

const (
	ApiMethodQuery  ApiMethod = "query"
	ApiMethodLogin  ApiMethod = "login"
	ApiMethodLogout ApiMethod = "logout"
)
