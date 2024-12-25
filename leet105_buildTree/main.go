package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder, inorder []int) *TreeNode {
	valToTree := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		valToTree[inorder[i]] = i
	}

	var build func([]int, int, int, []int, int, int) *TreeNode
	build = func(preorder []int, preStart int, preEnd int, inorder []int, inStart int, inEnd int) *TreeNode {
		if preStart > preEnd {
			return nil
		}
		//前序遍历的第一个元素就是 root节点
		rootVal := preorder[preStart]
		// rootVal在中序遍历中找到前后索引
		index := valToTree[rootVal]
		leftSize := index - inStart

		//先构造出当前节点
		root := &TreeNode{Val: rootVal}
		//递归左右子树
		root.Left = build(preorder, preStart+1, preStart+leftSize, inorder, inStart, index-1)
		root.Right = build(preorder, preStart+leftSize+1, preEnd, inorder, index+1, inEnd)
		return root
	}
	return build(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}
