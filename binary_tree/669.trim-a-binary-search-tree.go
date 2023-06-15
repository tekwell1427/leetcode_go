package main

/*
 * @lc app=leetcode id=669 lang=golang
 *
 * [669] Trim a Binary Search Tree
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	var helper func(node *TreeNode) *TreeNode
	helper = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		if node.Val >= low && node.Val <= high {
			node.Left = helper(node.Left)
			node.Right = helper(node.Right)
			return node
		}
		if node.Val < low {
			node = helper(node.Right)
		} else if node.Val > high {
			node = helper(node.Left)
		}
		return node
	}
	return helper(root)
}

// @lc code=end
