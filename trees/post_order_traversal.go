package main

func postorderTraversal(root *Tree) []int {
	nums := make([]int, 0)
	return getElements(root, nums)
}

func getElements(root *Tree, nums []int) []int {
	if root == nil {
		return nums
	}

	nums = getElements(root.Left, nums)
	nums = getElements(root.Right, nums)
	nums = append(nums, root.Val)

	return nums
}
