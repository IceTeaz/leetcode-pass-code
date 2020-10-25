package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	switch {
	case root.Val > val:
		return searchBST(root.Left, val)
	case root.Val < val:
		return searchBST(root.Right, val)
	default:
		return root
	}
}
