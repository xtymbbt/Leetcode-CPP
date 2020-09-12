package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var list []int

func postorderTraversal(root *TreeNode) []int {
	list = make([]int, 0)
	return dfs(root)
}

func dfs(root *TreeNode) []int {
	if root!=nil {
		dfs(root.Left)
		dfs(root.Right)
		list = append(list, root.Val)
	}
	return list
}
