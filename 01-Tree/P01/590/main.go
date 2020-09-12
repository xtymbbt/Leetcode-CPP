package main

type Node struct {
	Val int
	Children []*Node
}

var result []int

func postorder(root *Node) []int {
	result = make([]int, 0)
	if root != nil {
		dfs([]*Node{root})
	}
	return result
}

func dfs(children []*Node) {
	if len(children) != 0 {
		for _, child := range children {
			if child.Children != nil {
				dfs(child.Children)
			}
			result = append(result, child.Val)
		}
	}
}
