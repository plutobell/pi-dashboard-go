package main

import (
	"net/http"
	"testing"
)

func Test_Popen(t *testing.T) {
	if res, err := Popen("uptime"); res != "False" && err == nil {
		t.Log("Pass")
	} else {
		t.Error("Fail")
	}
}

func Benchmark_TimeConsumingFunction(b *testing.B) {
	for i := 0; i < 10000; i++ {
		Device()
	}
}

func Test_resolveTime(t *testing.T) {
	if uptime := resolveTime("1000000"); uptime == "11 days 13:46" {
		t.Log("Pass")
	} else {
		t.Error("Fail")
	}
}

func Test_bytesRound(t *testing.T) {
	if last := bytesRound(1073741824, 2); last == "1.0GB" {
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
