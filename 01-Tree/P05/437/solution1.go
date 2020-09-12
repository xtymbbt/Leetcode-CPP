package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	var l33 = &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	var rx2 = &TreeNode{
		Val:   -2,
		Left:  nil,
		Right: nil,
	}
	var r11 = &TreeNode{
		Val:   11,
		Left:  nil,
		Right: nil,
	}
	var rx3 = &TreeNode{
		Val:   -3,
		Left:  nil,
		Right: r11,
	}
	var r1 = &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	var r2 = &TreeNode{
		Val:   2,
		Left:  nil,
		Right: r1,
	}
	var l3 = &TreeNode{
		Val:   3,
		Left:  l33,
		Right: rx2,
	}
	var l5 = &TreeNode{
		Val:   5,
		Left:  l3,
		Right: r2,
	}
	var root = &TreeNode{
		Val:   10,
		Left:  l5,
		Right: rx3,
	}

	fmt.Println(pathSum(root, 8))
}

var pathNum int

func pathSum(root *TreeNode, sum int) int {
	pathNum = 0
	if root == nil {
		return 0
	}

	nodeAsRootFind(root, sum)
	dfs(root.Left, sum)
	dfs(root.Right, sum)

	return pathNum

}

func dfs(root *TreeNode, sum int) {
	if root == nil {
	} else {
		nodeAsRootFind(root, sum)
		dfs(root.Left, sum)
		dfs(root.Right, sum)
	}
}

func nodeAsRootFind(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	var val = root.Val
	if val == sum{
		pathNum++
	}

	return dfs2(root.Left, val, sum) || dfs2(root.Right, val, sum)
}

func dfs2(node *TreeNode, sum int, target int) bool {
	if node == nil {
		return false
	}
	sum += node.Val
	if sum == target{
		pathNum++
	}

	return dfs2(node.Left, sum, target) || dfs2(node.Right, sum, target)
}