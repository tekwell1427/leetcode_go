package main

/*
 * @lc app=leetcode id=337 lang=golang
 *
 * [337] House Robber III
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
	// first: steal current house, second: steal without current house
	var helper func(node *TreeNode) [2]int
	helper = func(node *TreeNode) [2]int {
		if node == nil {
			return [2]int{0, 0}
		}
		if node.Left == nil && node.Right == nil {
			return [2]int{node.Val, 0}
		}
		left := helper(node.Left)
		right := helper(node.Right)
		// steal with current house:
		withCur := left[1] + right[1] + node.Val
		withoutCur := max(left[0], left[1]) + max(right[0], right[1])

		return [2]int{withCur, withoutCur}
	}
	output := helper(root)
	return max(output[0], output[1])
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
