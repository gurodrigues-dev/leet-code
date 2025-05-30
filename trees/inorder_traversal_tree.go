package main

type Tree struct {
	Val   int
	Left  *Tree
	Right *Tree
}

func inorderTraversal(root *Tree) []int {
	nums := make([]int, 0)
	return inorder(root, nums)
}

func inorder(root *Tree, nums []int) []int {
	if root == nil {
		return nums
	}

	nums = inorder(root.Left, nums)
	nums = append(nums, root.Val)
	nums = inorder(root.Right, nums)

	return nums
}
