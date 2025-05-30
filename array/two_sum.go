package main

func twoSum(nums []int, target int) []int { // return index
	mapp := make(map[int]int)

	for index, value := range nums {
		if value, exists := mapp[target-value]; exists {
			return []int{value, index}
		}
		mapp[value] = index
	}

	return []int{}
}
