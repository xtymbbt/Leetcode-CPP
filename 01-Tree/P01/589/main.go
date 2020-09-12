package main

type Node struct {
	Val int
	Children []*Node
}

func preorder(root *Node) []int {
	var result = make([]int, 0)
	if root != nil {
		result = preorder2(root, result)
	}
	return result
}

func preorder2(root *Node, result []int) []int {
	if root != nil {
		result = append(result, root.Val)
		if root.Children != nil {
			for _, child := range root.Children {
				result = preorder2(child, result)
			}
		}
	}
	return result
}
