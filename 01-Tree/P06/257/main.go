package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []string

func binaryTreePaths(root *TreeNode) []string {
    res = make([]string, 0, 0)
    if root == nil {
        return res
    }
    var oneRes string
    oneRes += string(strconv.Itoa(root.Val))

    if root.Left == nil && root.Right == nil {
        res = append(res, oneRes)
        return res
    }
    oneRes += "->"

    dfs(root.Left, oneRes)
    dfs(root.Right, oneRes)

    return res
}

func dfs(root *TreeNode, oneRes string) {
    if root == nil {
    } else {
        oneRes += string(strconv.Itoa(root.Val))
        if root.Left == nil && root.Right == nil {
            res = append(res, oneRes)
        } else {
            oneRes += "->"
            dfs(root.Left, oneRes)
            dfs(root.Right, oneRes)
        }
    }
}