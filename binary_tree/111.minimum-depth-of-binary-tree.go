package main

import "math"

/*
 * @lc app=leetcode id=111 lang=golang
 *
 * [111] Minimum Depth of Binary Tree
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
func minDepth(root *TreeNode) int {
	minD := math.MaxInt16
	var helper func(node *TreeNode, depth int)
	helper = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		depth++
		if node.Left == nil && node.Right == nil {
			minD = min(minD, depth)
			return
		}
		helper(node.Left, depth)
		helper(node.Right, depth)
	}
	if root == nil {
		return 0
	}
	helper(root, 0)
	return minD
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
