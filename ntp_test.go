package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCount(t *testing.T) {
	host := "time.apple.com"
	time := ntpGetTime(host)

	gotType := reflect.TypeOf(time)

	if fmt.Sprint(gotType) == "time" {
		t.Fatalf("bad count for %s: got %d expected time.Time", host, gotType)
	}
}
