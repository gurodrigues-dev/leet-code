package main

/*
	1, 2, 3, 4, 5
*/

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func binSearch(nums []int, target, left, right int) int {
	for left < right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			right = mid
		}

		if nums[left] < target {
			left = mid + 1
		}
	}

	return -1
}

func exponentialSearch(nums []int, target int) int {
	if nums[0] == target {
		return 0
	}

	n, i := len(nums), 1

	for i < n && nums[i] < target {
		i *= 2
	}

	if nums[i] == target {
		return i
	}

	return binSearch(nums, target, i/2, min(i, n-1))
}

func main() {
	println(exponentialSearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 15))
}
