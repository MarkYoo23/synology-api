package synology_api

import (
	"encoding/json"
	"testing"
)

func TestClient_FileShare(t *testing.T) {
	var client = NewClient("...", "...")
	err := client.Login("...", "...")
	if err != nil {
		t.Error(err)
		return
	}

	defer client.Logout()

	share, err := client.GetShareFolder()
	if err != nil {
		t.Error(err)
		return
	}

	folder, err := client.GetFolder(share.Data.Shares[0].Path)
	if err != nil {
		t.Error(err)
		return
	}

	folder2, err := client.GetFolder(folder.Data.Files[1].Path)
	if err != nil {
		t.Error(err)
		return
	}

	indent, _ := json.MarshalIndent(folder2, "", "    ")
	t.Logf("Folder: %s", string(indent))
}
