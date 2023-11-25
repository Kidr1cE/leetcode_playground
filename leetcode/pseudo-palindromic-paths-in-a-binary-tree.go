/*
Date: 2023-11-25
ProblemID: 1457
ProblemName: 二叉树中的伪回文路径
*/
package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pseudoPalindromicPaths(root *TreeNode) int {
	var arr [9]int
	return dfs(root, arr)
}

func dfs(node *TreeNode, arr [9]int) int {
	if node == nil {
		return 0
	}
	if node.Left == nil && node.Right == nil {
		arr[node.Val-1]++
		key := 0
		for _, val := range arr {
			if val%2 != 0 {
				key++
			}
		}
		if key < 2 {
			return 1
		} else {
			return 0
		}
	} else {
		arr[node.Val-1]++
		return dfs(node.Left, arr) + dfs(node.Right, arr)
	}
}
