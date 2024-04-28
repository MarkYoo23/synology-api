package request

type ApiParam string

const (
	ApiParamQuery ApiParam = "query"

	// ApiParamAccount Used in login
	ApiParamAccount ApiParam = "account"
	ApiParamPasswd  ApiParam = "passwd"
	ApiParamSession ApiParam = "session"
	ApiParamFormat  ApiParam = "format"
)
