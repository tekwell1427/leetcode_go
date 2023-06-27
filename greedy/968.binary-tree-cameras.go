package main

/*
 * @lc app=leetcode id=968 lang=golang
 *
 * [968] Binary Tree Cameras
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 0: empty
// 1:
// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minCameraCover(root *TreeNode) int {
	output := 0
	var helper func(node *TreeNode) int
	helper = func(node *TreeNode) int {
		if node == nil {
			return 1
		}
		right := helper(node.Right)
		left := helper(node.Left)

		var signal int
		if right == 0 || left == 0 {
			output++
			signal = 2
		} else if (right == 2 && left >= 1) || (right >= 1 && left == 2) {
			signal = 1
		} else {
			signal = 0
		}
		return signal
	}

	if helper(root) == 0 {
		output++
	}
	return output
}

// @lc code=end
