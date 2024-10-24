package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {

	res := 0

	depth := 0

	var traverse func(node *TreeNode)

	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		depth++

		res = max(depth, res)

		traverse(node.Left)
		traverse(node.Right)
		depth--

	}

	traverse(root)

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
