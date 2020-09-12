package main

type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

func minDepth(root *TreeNode) int {
	var depth = 0
	var firstTime = true

	if root == nil {
		return 0
	}
	depth++
	var currentDepth = depth
	depth, firstTime = dfs(root.Left, currentDepth, depth, firstTime)
	depth, firstTime = dfs(root.Right, currentDepth, depth, firstTime)
	return depth
}

func dfs(root *TreeNode, currentDepth int, depth int, firstTime bool) (int, bool){
	if root == nil {
	} else {
		if root.Left == nil && root.Right == nil {
			currentDepth++
			if firstTime {
				depth = currentDepth
				firstTime = false
			} else if currentDepth < depth {
				depth = currentDepth
			}
		} else {
			currentDepth++
			depth, firstTime = dfs(root.Left, currentDepth, depth, firstTime)
			depth, firstTime = dfs(root.Right, currentDepth, depth, firstTime)
		}
	}
	return depth, firstTime
}