package main

import (
	"testing"
)

func Test_Popen(t *testing.T) {
	if res := Popen("uptime"); res != "False" {
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
