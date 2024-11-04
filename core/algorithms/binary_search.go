package algorithms

func BinarySearch(arr []int, target int) int {
	// We define lower and upper bounds to start the search
	lower_bound := 0
	upper_bound := len(arr) - 1

	// We will inspect the array until the lower bound is less than or equal to the upper bound
	for lower_bound <= upper_bound {

		// We calculate the middle index of the array. The result is truncated to an integer
		mid := (lower_bound + upper_bound) / 2

		// If the middle element is equal to the target, we return the index
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			// If the middle element is less than the target, we update the lower bound
			lower_bound = mid + 1
		} else {
			// If the middle element is greater than the target, we update the upper bound
			upper_bound = mid - 1
		}
	}

	// If the target is not found, we return -1
	return -1
}
