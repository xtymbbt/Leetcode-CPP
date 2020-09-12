package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}


func maxDepth(root *TreeNode) int {
	var depth = 0
	if root == nil {
		return depth
	}
	depth++
	var hereDepth = depth
	depth = dfs(root.Left, hereDepth, depth)
	depth = dfs(root.Right, hereDepth, depth)
	return depth
}

func dfs(root *TreeNode, hereDepth int, depth int)  int {
	if root == nil {
		if hereDepth > depth {
			depth = hereDepth
		}
	} else {
		hereDepth++
		depth = dfs(root.Left, hereDepth, depth)
		depth = dfs(root.Right, hereDepth, depth)
	}
	return depth
}
