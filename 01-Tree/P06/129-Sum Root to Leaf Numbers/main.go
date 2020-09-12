package main

import "math"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var lists [][]int

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lists = make([][]int, 0, 0)
	var list = []int{root.Val}
	if root.Left == nil && root.Right == nil {
		var x = make([]int, len(list))
		copy(x, list)
		lists = append(lists, x)
	} else {
		dfs(root.Left, list)
		dfs(root.Right, list)
	}
	return calculate(lists)
}

func dfs(root *TreeNode, list []int) {
	if root == nil {
	} else {
		list = append(list, root.Val)
		if root.Left == nil && root.Right == nil {
			var x = make([]int, len(list))
			copy(x, list)
			lists = append(lists, x)
		} else {
			dfs(root.Left, list)
			dfs(root.Right, list)
		}
	}
}

func calculate(lists [][]int) int {
	var sum = 0
	for _, list := range lists {
		var len = len(list)-1
		for _, val := range list {
			sum += val * int(math.Pow10(len))
			len--
		}
	}
	return sum
}