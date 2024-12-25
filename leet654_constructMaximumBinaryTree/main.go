package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return build(nums, 0, len(nums)-1)
}

func build(nums []int, l, r int) *TreeNode {
	if l < r {
		return nil
	}

	index := -1
	maxVal := -1 << 31
	for i := l; i < r; i++ {
		if maxVal < nums[index] {
			index = i
			maxVal = nums[index]
		}
	}

	root := &TreeNode{Val: maxVal}
	root.Left = build(nums, l, index-1)
	root.Right = build(nums, index+1, r)
	return root
}
