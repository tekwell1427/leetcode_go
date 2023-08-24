package main

/*
 * @lc app=leetcode id=112 lang=golang
 *
 * [112] Path Sum
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
func hasPathSum(root *TreeNode, targetSum int) bool {
	var helper func(node *TreeNode, curSum int) bool
	helper = func(node *TreeNode, curSum int) bool {
		if node == nil {
			return false
		}
		curSum += node.Val
		if node.Left == nil && node.Right == nil {
			if curSum == targetSum {
				return true
			}
			return false
		}
		if helper(node.Left, curSum) == true || helper(node.Right, curSum) == true {
			return true
		}
		return false
	}
	return helper(root, 0)
}

// @lc code=end
