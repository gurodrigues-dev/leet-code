package main

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)

	for _, num := range nums {
		if _, exists := m[num]; exists {
			return true
		}

		m[num] = true
	}

	return false
}
