package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var root1leaf = leaf(root1)
	var root2leaf = leaf(root2)
	if len(root1leaf) != len(root2leaf) {
		return false
	}
	for i, i2 := range root1leaf {
		if root2leaf[i] == i2 {
		} else {
			return false
		}
	}
	return true
}

func leaf(root *TreeNode) []int {
	var leaf = make([]int, 0, 0)
	if root == nil {
		return leaf
	}
	if root.Left == nil && root.Right == nil {
		leaf = append(leaf, root.Val)
	}
	leaf = dfs(root.Left, leaf)
	leaf = dfs(root.Right, leaf)
	return leaf
}

func dfs(root *TreeNode, leaf []int) []int {
	if root == nil {
		return leaf
	}
	if root.Left == nil && root.Right == nil {
		leaf = append(leaf, root.Val)
	}
	leaf = dfs(root.Left, leaf)
	leaf = dfs(root.Right, leaf)
	return leaf
}