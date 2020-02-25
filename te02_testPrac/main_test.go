package main

import (
	"testing"
)

func TestGetStatusCode(t *testing.T) {
	if getStatusCode("https://www.google.com") != 200 {
		t.Error("Google is not responding with a 200")
	}
}
