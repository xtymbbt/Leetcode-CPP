package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var res = make([][]int, 0, 0)
	if root == nil {
		return res
	}
	res = nextLevel([]*TreeNode{root}, res)
	return res
}

func nextLevel(levelNode []*TreeNode, res [][]int) [][]int {
	if len(levelNode) == 0 {
	} else {
		var levelRes = make([]int, 0, 0)
		var nextLevelNode = make([]*TreeNode, 0, 0)
		for _, node := range levelNode {
			if node != nil {
				levelRes = append(levelRes, node.Val)
				nextLevelNode = append(nextLevelNode, node.Left, node.Right)
			}
		}
		if len(levelRes) == 0 {

		} else {
			res = append(res, levelRes)
			res = nextLevel(nextLevelNode, res)
		}
	}
	return res
}