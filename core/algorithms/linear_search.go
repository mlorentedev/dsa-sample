package algorithms

func LinearSearch(arr []int, target int) int {
	// Iterate over each value in the array
	for i, val := range arr {
		// If the value is equal to the target, return the index
		if val == target {
			return i
		}
	}
	// If the target value is not found, return -1
	return -1
}
