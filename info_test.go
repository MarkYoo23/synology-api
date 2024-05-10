package synology_api

import (
	"encoding/json"
	"testing"
)

func TestRetrieveApiInfo(t *testing.T) {
	var client = NewClient("your_synology_server_name", "your_synology_url")
	info, err := client.GetApiInfo()
	if err != nil {
		t.Error(err)
		return
	}

	// json pretty print
	indent, err := json.MarshalIndent(info, "", "    ")
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("ApiInfo: %s", string(indent))
}
