package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return dfs(root.Left, root.Right)
}

func dfs(p *TreeNode, q *TreeNode) bool {
	if p != nil && q != nil {
		if p.Val != q.Val {return false}
		return dfs(p.Left, q.Right) && dfs(p.Right, q.Left)
	} else if (p != nil && q == nil) || (p == nil && q != nil) {
		return false
	} else {
		return true
	}
}
