package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	var val = root.Val
	if val == sum && root.Left == nil && root.Right == nil {
		return true
	}

	return dfs(root.Left, val, sum) || dfs(root.Right, val, sum)
}

func dfs(node *TreeNode, sum int, target int) bool {
	if node == nil {
		return false
	}
	sum += node.Val
	if sum == target && node.Left == nil && node.Right == nil {
		return true
	}

	return dfs(node.Left, sum, target) || dfs(node.Right, sum, target)
}