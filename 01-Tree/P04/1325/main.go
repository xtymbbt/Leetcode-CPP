package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	return deal(root, target)
}

func deal(node *TreeNode, target int) *TreeNode {
	if node == nil {
		return nil
	}
	node.Left = deal(node.Left, target)
	node.Right = deal(node.Right, target)
	if node.Left == nil && node.Right == nil && node.Val == target {
		return nil
	}
	return node
}