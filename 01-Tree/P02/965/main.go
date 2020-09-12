package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var value = root.Val
	return dfs(root.Left, value) && dfs(root.Right, value)
}

func dfs(root *TreeNode, value int) bool {
	if root == nil {
		return true
	}
	if root.Val == value {
		return dfs(root.Left, value) && dfs(root.Right, value)
	} else {
		return false
	}
}