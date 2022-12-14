package main

import (
	"reflect"
	"testing"
)

func TestHello(t *testing.T) {
	want := "Hello, my dude!"
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func TestSeparateLinesInFileByElf(t *testing.T) {
	want := [][]int{{1000, 2000, 3000}, {4000}, {5000, 6000}, {7000, 8000, 9000}, {10000}}
	got := separateLinesInFileByElf("aoc22_in_test.txt")
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Separated test file = %v, want %v", got, want)
	}
}

func TestSumOverElfs(t *testing.T) {
	want := []int{6000, 4000, 11000, 24000, 10000}
	elfs := separateLinesInFileByElf("aoc22_in_test.txt")
	got := sumOverElfs(elfs)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("%q not equal to %q", got, want)
	}
}
