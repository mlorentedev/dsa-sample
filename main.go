package main

import (
	"fmt"
	"github.com/manulorente/dsa-sample/core"
	"time"
)

func main() {
	dataSample := []int{1, 2, 3, 4, 5}
	targetToFind := 3
	startTime := time.Now()
	result := algorithms.LinearSearch(dataSample, targetToFind)
	elapsedTime := time.Since(startTime)
	fmt.Printf("Result: %d\n", result)
	fmt.Printf("Elapsed time: %s\n", elapsedTime)
}
