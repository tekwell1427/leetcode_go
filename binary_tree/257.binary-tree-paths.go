package main

import (
	"strconv"
)

/*
 * @lc app=leetcode id=257 lang=golang
 *
 * [257] Binary Tree Paths
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
func binaryTreePaths(root *TreeNode) []string {
	output := []string{}
	var helper func(node *TreeNode, subString string)
	helper = func(node *TreeNode, subString string) {
		if node == nil {
			return
		}
		if node != root {
			subString += "->" + strconv.Itoa(node.Val)
		}
		if node.Left == nil && node.Right == nil {
			output = append(output, subString)
			return
		}
		helper(node.Left, subString)
		helper(node.Right, subString)
	}
	helper(root, strconv.Itoa(root.Val))
	return output
}

// @lc code=end
