package tests

import (
	"math/rand"
	"testing"

	algorithms "github.com/manulorente/dsa-sample/core"
)

const (
	size = 100000
	seed = 33
)

func generateSortedArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = i
	}
	return arr
}

func generateUnorderedArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(size)
	}
	return arr
}

func BenchmarkLinearSearchSorted(b *testing.B) {

	arr := generateSortedArray(size)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		algorithms.LinearSearch(arr, size-1)
	}
}

func BenchmarkLinearSearchUnordered(b *testing.B) {

	arr := generateUnorderedArray(size)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		algorithms.LinearSearch(arr, size-1)
	}
}

func BenchmarkBinarySearchSorted(b *testing.B) {

	arr := generateSortedArray(size)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		algorithms.BinarySearch(arr, size-1)
	}
}
