package tests

import (
	"testing"

	algorithms "github.com/mlorentedev/dsa-sample/core/algorithms"
)

func TestBinarySearch(t *testing.T) {

	tests := []struct {
		arr      []int
		target   int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, 1},
		{[]int{1, 2, 3, 4, 5}, 4, 3},
		{[]int{1, 2, 3, 4, 5}, 6, -1},
	}

	for _, test := range tests {
		actual := algorithms.BinarySearch(test.arr, test.target)
		if actual != test.expected {
			t.Errorf("Expected %d, but got %d", test.expected, actual)
		}
	}
}
