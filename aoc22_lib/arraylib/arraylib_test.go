package arraylib

import (
	"testing"
)

func TestGetLargestElementAndIndex(t *testing.T) {
	arr_of_ints := []int{6000, 4000, 11000, 24000, 10000}
	want_index := 3
	want_val := 24000
	got_index, got_value := GetLargestElementAndIndex(arr_of_ints)
	if want_index != got_index && want_val != got_value {
		t.Errorf("Computed Index, Value (%v, %v) not equal to Given Index, Value: (%v, %v)", got_index, got_value, want_index, want_val)
	}
}

func TestSum(t *testing.T) {
	var want int = 42
	array_to_sum := []int{2, 30, 10, 0, 0, 0}
	got := Sum(array_to_sum)
	if got != want {
		t.Errorf("Got %v not equal Want %v", got, want)
	}
}
