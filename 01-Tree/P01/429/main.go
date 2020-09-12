package main

type Node struct {
	Val int
	Children []*Node
}

var result [][]int
var single []int
var childrenNodes []*Node

func levelOrder(root *Node) [][]int {
	result = make([][]int, 0)
	if root == nil {
		return result
	}
	inSameLevel([]*Node{root})
	return result
}

func inSameLevel(nodes []*Node) {
	if len(nodes) != 0 {
		single = make([]int, 0)
		childrenNodes = make([]*Node, 0)
		for _, node := range nodes {
			single = append(single, node.Val)
			if node.Children != nil {
				childrenNodes = append(childrenNodes, node.Children...)
			}
		}
		result = append(result, single)
		inSameLevel(childrenNodes)
	}
}