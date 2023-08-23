package main

/*
 * @lc app=leetcode id=101 lang=golang
 *
 * [101] Symmetric Tree
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
func isSymmetric(root *TreeNode) bool {
	var helper func(left, right *TreeNode) bool
	helper = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if (left == nil && right != nil) || (left != nil && right == nil) {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		return helper(left.Left, right.Right) && helper(left.Right, right.Left)
	}
	return helper(root.Left, root.Right)
}

// @lc code=end
