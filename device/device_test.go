package device

import (
	"reflect"
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
		Info()
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

func Test_struct2Map(t *testing.T) {
	host := new(Host)
	host.Get()
	hostMap, _ := struct2Map(host, "json")
	if typeOf := reflect.TypeOf(hostMap); typeOf.Kind() == reflect.Map {
		t.Log("Pass")
	} else {
		t.Error("Fail")
	}
}

func Test_mergeMap(t *testing.T) {
	map1 := make(map[string]interface{})
	map2 := map[string]interface{}{"a": "Apple"}
	if map1 := mergeMap(map1, map2); map1["a"] == "Apple" {
		t.Log("Pass")
	} else {
		t.Error("Fail")
	}
}
