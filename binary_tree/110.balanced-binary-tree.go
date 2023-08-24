package main

/*
 * @lc app=leetcode id=110 lang=golang
 *
 * [110] Balanced Binary Tree
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
func isBalanced(root *TreeNode) bool {
	var helper func(node *TreeNode, depth int) int
	helper = func(node *TreeNode, depth int) int {
		if node == nil {
			return depth
		}
		depth++
		left, right := helper(node.Left, depth), helper(node.Right, depth)
		if left == -1 || right == -1 {
			return -1
		}
		if abs(left-right) > 1 {
			return -1
		}
		return max(left, right)
	}
	return helper(root, 0) != -1
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
