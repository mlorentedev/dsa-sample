package problems

func hasDuplicateValue(arr []int) bool {
	existingValues := make(map[int]bool)
	for _, value := range arr {
		if existingValues[value] {
			return true
		}
		existingValues[value] = true
	}
	return false
}
