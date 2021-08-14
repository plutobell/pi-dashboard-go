package server

import (
	"net/http"
	"testing"
)

func Test_getNowUsernameAndPassword(t *testing.T) {
	if username, password := getNowUsernameAndPassword(); username+":"+password == Auth {
		t.Log("Pass")
	} else {
		t.Error("Fail")
	}
}

func Test_getFileSystem(t *testing.T) {
	if _, ok := getFileSystem(false).(http.FileSystem); ok {
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
	if nowVersion, downloadURL := getLatestVersionFromGitHub(); nowVersion != "" && len(downloadURL) > 0 {
		t.Log("Pass")
	} else {
		t.Error("Fail")
	}
}
