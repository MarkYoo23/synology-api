package request

type ApiMethod string

const (
	ApiMethodQuery           ApiMethod = "query"
	ApiMethodLogin           ApiMethod = "login"
	ApiMethodLogout          ApiMethod = "logout"
	ApiMethodShareFolderList ApiMethod = "list_share"
	ApiMethodFolderList      ApiMethod = "list"
	ApiMethodFileDownload    ApiMethod = "download"
)
