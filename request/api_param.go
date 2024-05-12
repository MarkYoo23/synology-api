package request

type ApiParam string

const (
	ApiParamQuery ApiParam = "query"

	ApiParamAccount ApiParam = "account"
	ApiParamPasswd  ApiParam = "passwd"
	ApiParamSession ApiParam = "session"
	ApiParamFormat  ApiParam = "format"

	ApiParamFolderPath       ApiParam = "folder_path"
	ApiParamFolderAdditional ApiParam = "additional"

	ApiParamDownloadPath ApiParam = "path"
	ApiParamDownloadMode ApiParam = "mode"
)
