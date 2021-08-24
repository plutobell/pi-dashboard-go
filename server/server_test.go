package server

import (
	"net/http"
	"testing"

	"github.com/plutobell/pi-dashboard-go/config"
)

func Test_getNowUsernameAndPassword(t *testing.T) {
	if username, password := getNowUsernameAndPassword(); username+":"+password == config.Auth {
		t.Log("Pass")
	} else {
		t.Error("Fail")
	}
}

func Test_getFileSystem(t *testing.T) {
	if _, ok := getFileSystem(false, "btns").(http.FileSystem); ok {
		t.Log("Pass")
	} else {
		t.Error("Fail")
	}
}

func Test_getRandomString(t *testing.T) {
	if res := getRandomString(16); len(res) == 16 {
		t.Log("Pass")
	} else {
		t.Error("Fail")
	}
}

func Test_getLatestVersionFromGitHub(t *testing.T) {
	if nowVersion, _, downloadURL := getLatestVersionFromGitHub(); nowVersion != "" && len(downloadURL) > 0 {
		t.Log("Pass")
	} else {
		t.Error("Fail")
	}
}
