package algorithms

// LinearSearch searches for a given element in a list of elements
func LinearSearch(arr []int, target int) int {
	for i, val := range arr {
		if val == target {
			return i
		}
	}
	return -1
}
