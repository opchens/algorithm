package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	maxdia := 0

	var maxDepth func(node *TreeNode) int
	maxDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftMax := maxDepth(node.Left)
		rightMax := maxDepth(node.Right)

		depth := leftMax + rightMax
		maxdia = max(maxdia, depth)

		return max(leftMax, rightMax) + 1
	}
	maxDepth(root)
	return maxdia
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
