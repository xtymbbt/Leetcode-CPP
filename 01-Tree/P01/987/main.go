package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type Position struct {
	x int
	y int
	Val int
}

var res []Position

func inorderTraversal(root *TreeNode) []Position {
	res = make([]Position, 0)
	dfs(root, 0, 0)
	return res
}

func dfs(root *TreeNode, x int, y int) {
	if root != nil {
		dfs(root.Left, x - 1, y - 1)
		res = append(res, Position{
			x:   x,
			y:   y,
			Val: root.Val,
		})
		dfs(root.Right, x + 1, y - 1)
	}
}

func verticalTraversal(root *TreeNode) [][]int {
	inorderTraversal(root)
	sortRes()
	finalResPos := sortAllPart()
	finalRes := make([][]int, 0, 0)
	for _, po := range finalResPos {
		oneRes := make([]int, 0, 0)
		for _, position := range po {
			oneRes = append(oneRes, position.Val)
		}
		finalRes = append(finalRes, oneRes)
	}
	return finalRes
}

func sortRes() {
	n := len(res)
	for i := 1; i < n; i++ {
		tmp := res[i]
		j := i - 1
		for j >= 0 && res[j].x > tmp.x { //左边比右边大
			res[j+1] = res[j] //右移1位
			j--               //扫描前一个数
		}
		res[j+1] = tmp //添加到小于它的数的右边
	}
}

func sortAllPart() [][]Position {
	var finalRes = make([][]Position, 0, 0)
	var oneRes = make([]Position, 0, 0)
	for i, re := range res {
		if i == 0 {
			oneRes = append(oneRes, re)
		} else if i == len(res)-1 {
			if re.x == res[i-1].x {
				oneRes = append(oneRes, re)
				oneRes = sortTinyPart(oneRes)
				finalRes = append(finalRes, oneRes)
			} else {
				oneRes = sortTinyPart(oneRes)
				finalRes = append(finalRes, oneRes)
				oneRes = []Position{re}
				finalRes = append(finalRes, oneRes)
			}
		} else {
			if re.x == res[i-1].x {
				oneRes = append(oneRes, re)
			} else {
				oneRes = sortTinyPart(oneRes)
				finalRes = append(finalRes, oneRes)
				oneRes = []Position{re}
			}
		}
	}
	return finalRes
}

func sortTinyPart(nums []Position) []Position {
	n := len(nums)
	for i := 0; i < n - 1; i++ {
		exchange := false
		for j := 0; j < n-i-1; j++ {
			if nums[j].y < nums[j+1].y || (nums[j].y == nums[j+1].y && nums[j].Val > nums[j+1].Val){ //两两比较，前面比后面大
				nums[j], nums[j+1] = nums[j+1], nums[j] //交换
				exchange = true
			}
		}
		if !exchange {
			return nums
		}
	}
	return nums
}