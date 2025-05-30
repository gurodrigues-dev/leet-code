package main

func BinarySearch(nums []int, target int) int {
	left := 0
	right := len(nums)

	for left < right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			right = mid
		}

		if nums[mid] < target {
			left = mid + 1
		}
	}

	return -1
}
