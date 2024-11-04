package tests

import (
	"testing"

	algorithms "github.com/mlorentedev/dsa-sample/core/algorithms"
)

func TestLinearSearch(t *testing.T) {

	tests := []struct {
		arr      []int
		target   int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, 2},
		{[]int{1, 2, 3, 4, 5}, 6, -1},
	}

	for _, test := range tests {
		actual := algorithms.LinearSearch(test.arr, test.target)
		if actual != test.expected {
			t.Errorf("Expected %d, but got %d", test.expected, actual)
		}
	}
}
