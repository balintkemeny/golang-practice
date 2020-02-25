package main

import "testing"

func TestAddTwo(t *testing.T) {
	if addTwo(3) != 5 {
		t.Error("Expected 5, got something else")
	}
}
