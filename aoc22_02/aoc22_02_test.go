package main

import (
	"testing"
)

func TestNothing(t *testing.T) {
	got := 1
	want := 1
	nothing()
	if got != want {
		t.Errorf("TEST FAILED!")
	}
}
