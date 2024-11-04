package algorithms

func BubbleSort(arr []int) []int {
	// We define a variable to keep track of whether we swapped any elements in the last iteration
	swapped := true

	// We will continue to iterate through the array until no elements are swapped
	for swapped {
		// We set the swapped variable to false at the beginning of each iteration
		swapped = false

		// We iterate through the array, comparing each element with the next one
		for i := 0; i < len(arr)-1; i++ {
			// If the current element is greater than the next element, we swap them
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				// We set the swapped variable to true to indicate that we swapped elements
				swapped = true
			}
		}
	}

	// We return the sorted array
	return arr
}
