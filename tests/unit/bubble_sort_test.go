package tests

import (
	"reflect"
	"testing"

	algorithms "github.com/mlorentedev/dsa-sample/core/algorithms"
)

func TestBubbleSort(t *testing.T) {

	tests := []struct {
		arr      []int
		expected []int
	}{
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 3, 2, 5, 4}, []int{1, 2, 3, 4, 5}},
	}

	for _, test := range tests {
		actual := algorithms.BubbleSort(test.arr)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("Expected %v, but got %v", test.expected, actual)
		}
	}
}
