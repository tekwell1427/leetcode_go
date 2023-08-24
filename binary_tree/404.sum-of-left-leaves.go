package main

/*
 * @lc app=leetcode id=404 lang=golang
 *
 * [404] Sum of Left Leaves
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
func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	var helper func(node *TreeNode, isLeft bool)
	helper = func(node *TreeNode, isLeft bool) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			if isLeft {
				sum += node.Val
			}
			return
		}
		helper(node.Left, true)
		helper(node.Right, false)
	}
	helper(root, false)
	return sum
}

// @lc code=end
