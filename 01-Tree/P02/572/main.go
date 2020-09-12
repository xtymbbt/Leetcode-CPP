package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	var send = &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	var s4 = &TreeNode{
		Val:   1,
		Left:  send,
		Right: nil,
	}
	var s3 = &TreeNode{
		Val:   1,
		Left:  s4,
		Right: nil,
	}
	var s2 = &TreeNode{
		Val:   1,
		Left:  s3,
		Right: nil,
	}
	var s1 = &TreeNode{
		Val:   1,
		Left:  s2,
		Right: nil,
	}
	var s = &TreeNode{
		Val:   1,
		Left:  nil,
		Right: s1,
	}
	var t1 = &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	var t = &TreeNode{
		Val:   1,
		Left:  t1,
		Right: nil,
	}
	fmt.Println(isSubtree(s, t))
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	return check(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func check(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Val == b.Val {
		return check(a.Left, b.Left) && check(a.Right, b.Right)
	}
	return false
}