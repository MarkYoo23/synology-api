package client

import "testing"

func TestClient_LoginWithLogout(t *testing.T) {
	var client = NewClient("your_synology_server_name", "your_synology_url")
	err := client.Login("your_id", "your_password")
	if err != nil {
		t.Error(err)
		return
	}

	if client.IsLogin() {
		t.Logf("Login Success: %s", client.deviceId)
	} else {
		t.Error("Login Failed")
		return
	}

	err = client.Logout()
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("Logout Success")
}
