package main

import (
	"testing"
)

func Test_Popen(t *testing.T) {
	if res := Popen("uptime"); res != "False" {
		t.Error("Fail")
	} else {
		t.Log("Pass")
	}
}

func Benchmark_TimeConsumingFunction(b *testing.B) {
	Device()
}
