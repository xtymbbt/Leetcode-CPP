package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
solution 1: slowly way.
 */
func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return dfs(root.Left) && dfs(root.Right) && math.Abs(float64(depth(root.Left)-depth(root.Right))) <= 1
}

func dfs(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return dfs(root.Left) && dfs(root.Right) && math.Abs(float64(depth(root.Left)-depth(root.Right))) <= 1
}

func depth(root *TreeNode) int {
    var depth = 0
    if root == nil {
        return depth
    }
    depth++
    var hereDepth = depth
    depth = dfs2(root.Left, hereDepth, depth)
    depth = dfs2(root.Right, hereDepth, depth)
    return depth
}

func dfs2(root *TreeNode, hereDepth int, depth int)  int {
    if root == nil {
        if hereDepth > depth {
            depth = hereDepth
        }
    } else {
        hereDepth++
        depth = dfs2(root.Left, hereDepth, depth)
        depth = dfs2(root.Right, hereDepth, depth)
    }
    return depth
}

/**
solution 2: fast way:
 */

func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }


    lh := height(root.Left)
    rh := height(root.Right)

    if abs(lh-rh) > 1 {
        return false
    }

    return isBalanced(root.Left) && isBalanced(root.Right)
}

func height(root *TreeNode)int{
    if root == nil {
        return 0
    }

    l := height(root.Left)
    r := height(root.Right)
    if l > r {
        return l+1
    }else{
        return r+1
    }
}

func abs(n int) int{
    if n < 0 {
        return -n
    }
    return n
}