package main

import "math"

/*
 * @lc app=leetcode id=501 lang=golang
 *
 * [501] Find Mode in Binary Search Tree
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
func findMode(root *TreeNode) []int {
	output := []int{}
	preNode := &TreeNode{Val: math.MinInt32}
	var maxValCount, curValCount int
	var helper func(node *TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}
		helper(node.Left)

		if preNode.Val == node.Val {
			curValCount++
		} else {
			curValCount = 1
		}
		if curValCount == maxValCount {
			output = append(output, node.Val)
		} else if curValCount > maxValCount {
			output = []int{node.Val}
			maxValCount = curValCount
		}

		preNode = node
		helper(node.Right)
	}
	helper(root)
	return output
}

// @lc code=end
