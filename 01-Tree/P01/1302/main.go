package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type Position struct {
	y int
	Val int
}

var list []Position

func preorderTraversal(root *TreeNode) []Position {
	list = make([]Position, 0)
	return dfs(root, 0)
}

func dfs(root *TreeNode, y int) []Position {
	if root!=nil {
		list = append(list, Position{
			y:   y,
			Val: root.Val,
		})
		dfs(root.Left, y - 1)
		dfs(root.Right, y - 1)
	}
	return list
}

func deepestLeavesSum(root *TreeNode) int {
	preorderTraversal(root)
	res := 0
	depth := 0
	for i := 0; i < len(list); i++ {
		if list[i].y == depth {
			res += list[i].Val
		} else if list[i].y < depth {
			depth = list[i].y
			res = 0
			res += list[i].Val
		}
	}
	return res
}

/**
solution 2:
 */
var max int
var sum int

func deepestLeavesSum2(root *TreeNode) int {
	max, sum = 0, 0
	dfs(root, 0)
	return sum
}
func dfs2(root *TreeNode, level int) {
	if root == nil {
		return
	}
	if max < level {
		max = level
		sum = root.Val
	} else if max == level {
		sum += root.Val
	}

	dfs(root.Left, level+1)
	dfs(root.Right, level+1)
}