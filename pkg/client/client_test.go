package client

import (
	"encoding/json"
	"testing"
)

var client = NewClient("http://192.168.0.101:5000")

func TestRetrieveApiInfo(t *testing.T) {
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
