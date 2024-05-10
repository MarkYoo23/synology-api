package request

type ApiCgi string

const (
	ApiCgiQuery     ApiCgi = "query.cgi"
	ApiCgiAuthV1to3 ApiCgi = "auth.cgi"
	ApiCgiEntry     ApiCgi = "entry.cgi"
)
