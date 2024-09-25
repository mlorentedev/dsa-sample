package tests

import (
	"github.com/manulorente/dsa-sample/core"
	"testing"
)

func BenchmarkLinearSearch(b *testing.B) {
	arr := make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		arr[i] = i
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		algorithms.LinearSearch(arr, 999999)
	}
}
