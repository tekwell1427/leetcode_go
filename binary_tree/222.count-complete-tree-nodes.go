package main

/*
 * @lc app=leetcode id=222 lang=golang
 *
 * [222] Count Complete Tree Nodes
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
func countNodes(root *TreeNode) int {
	var helper func(node *TreeNode) int
	helper = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftDepth := 0
		left := node.Left
		for left != nil {
			left = left.Left
			leftDepth++
		}
		rightDepth := 0
		right := node.Right
		for right != nil {
			right = right.Right
			rightDepth++
		}
		if leftDepth == rightDepth {
			return 1<<(leftDepth+1) - 1
		}
		return helper(node.Left) + helper(node.Right) + 1
	}
	return helper(root)
}

// @lc code=end
