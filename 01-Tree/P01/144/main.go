package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var list []int

func preorderTraversal(root *TreeNode) []int {
	list = make([]int, 0)
	return dfs(root)
}

func dfs(root *TreeNode) []int {
	if root!=nil {
		list = append(list, root.Val)
		dfs(root.Left)
		dfs(root.Right)
	}
	return list
}
