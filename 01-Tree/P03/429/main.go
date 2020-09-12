package main

type Node struct {
    Val int
    Children []*Node
}

func levelOrder(root *Node) [][]int {
	var res = make([][]int, 0, 0)

	if root == nil {
		return res
	}

	res = append(res, []int{root.Val})

	res = nextLevel(root.Children, res)
	return res
}

func nextLevel(nodes []*Node, res [][]int) [][]int{
	if len(nodes) == 0 {
	} else {
		var levelRes = make([]int, 0, 0)
		var children = make([]*Node, 0, 0)
		for _, node := range nodes {
			if node != nil {
				levelRes = append(levelRes, node.Val)
				children = append(children, node.Children...)
			}
		}
		if len(levelRes) != 0 {
			res = append(res, levelRes)
			res = nextLevel(children, res)
		}
	}
	return res
}
