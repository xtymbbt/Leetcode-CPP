package main

var counter int
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	vals := map[int]int{
		0: 1,
	}
	return dfs(root, sum, 0, vals)
}

func dfs(nd *TreeNode, sum int, cur int, vals map[int]int) int {
	if nd == nil {
		return 0
	}

	num := 0
	cur += nd.Val
	if val, found := vals[cur - sum]; found {
		num += val
	}

	if val, found := vals[cur]; found {
		vals[cur] = val + 1
	} else {
		vals[cur] = 1
	}

	num += dfs(nd.Left, sum, cur, vals)
	num += dfs(nd.Right, sum, cur, vals)

	if val, found := vals[cur]; found {
		vals[cur] = val - 1
	}

	return num
}
