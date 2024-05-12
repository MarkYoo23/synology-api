package synology_api

import (
	"io"
	"testing"
)

func TestClient_OpenFile(t *testing.T) {
	var client = NewClient("...", "...")
	err := client.Login("...", "...")
	if err != nil {
		return
	}

	defer client.Logout()

	rc, err := client.OpenFile("/.../.../sample01.mp4")
	if err != nil {
		return
	}

	bytes, _ := io.ReadAll(rc)
	t.Logf("File: %s", string(bytes))
}

func TestClient_DownloadFile(t *testing.T) {
	var client = NewClient("...", "...")
	err := client.Login("...", "...")
	if err != nil {
		return
	}

	defer client.Logout()

	rc, err := client.DownloadFile("/.../.../sample01.mp4")
	if err != nil {
		return
	}

	bytes, _ := io.ReadAll(rc)
	t.Logf("File: %s", string(bytes))
}
