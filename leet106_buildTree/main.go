package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	//存储索引映射
	valToIndex := make(map[int]int)
	for i, v := range inorder {
		valToIndex[v] = i
	}

	// build 函数的定义：
	// 后序遍历数组为 postorder[postStart..postEnd]，
	// 中序遍历数组为 inorder[inStart..inEnd]，
	// 构造二叉树，返回该二叉树的根节点
	var build func(inorder []int, postorder []int, inStart int, inEnd int, postStart int, postEnd int) *TreeNode
	build = func(inorder []int, postorder []int, inStart int, inEnd int, postStart int, postEnd int) *TreeNode {
		if inStart > inEnd {
			return nil
		}
		// root 节点对应的值就是后续便利数组的最后一个元素
		rootVal := postorder[postEnd]
		// rootVal 在中序遍历数组中的索引
		index := valToIndex[rootVal]

		// 左子树的节点个数
		leftSize := index - inStart
		root := &TreeNode{Val: rootVal}
		// 递归构造左右子树
		root.Left = build(inorder, postorder, inStart, index-1, postStart, postStart+leftSize-1)
		root.Right = build(inorder, postorder, index+1, inEnd, postStart+leftSize, postEnd-1)
		return root
	}

	return build(inorder, postorder, 0, len(inorder)-1, 0, len(postorder)-1)
}
