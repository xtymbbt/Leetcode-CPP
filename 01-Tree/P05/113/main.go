package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	var r1 = &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	var l5 = &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}

	var l13 = &TreeNode{
		Val:   13,
		Left:  nil,
		Right: nil,
	}
	
	var r4 = &TreeNode{
		Val:   4,
		Left:  l5,
		Right: r1,
	}
	
	var r8 = &TreeNode{
		Val:   8,
		Left:  l13,
		Right: r4,
	}

	var r2 = &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}

	var l7 = &TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}

	var l11 = &TreeNode{
		Val:   11,
		Left:  l7,
		Right: r2,
	}
	
	var l4 = &TreeNode{
		Val:   4,
		Left:  l11,
		Right: nil,
	}
	
	var root = &TreeNode{
		Val:   5,
		Left:  l4,
		Right: r8,
	}

	fmt.Println(pathSum(root, 22))
}

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return make([][]int, 0, 0)
	}
	var val = root.Val
	if val == sum && root.Left == nil && root.Right == nil {
		return [][]int{[]int{root.Val}}
	}

	var finalRes = make([][]int, 0, 0)
	var res = []int{root.Val}
	finalRes = dfs(root.Left, val, sum, res, finalRes)
	finalRes = dfs(root.Right, val, sum, res, finalRes)

	return finalRes
}

func dfs(root *TreeNode, val int, sum int, res []int, finalRes [][]int) [][]int {
	if root == nil {
		return finalRes
	}
	val += root.Val
	res = append(res, root.Val)
	if val == sum && root.Left == nil && root.Right == nil {
		var x = make([]int, len(res))
		copy(x, res)
		finalRes = append(finalRes, x)
		return finalRes
	}

	finalRes = dfs(root.Left, val, sum, res, finalRes)
	finalRes = dfs(root.Right, val, sum, res, finalRes)

	return finalRes
}
