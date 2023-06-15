package main

/*
 * @lc app=leetcode id=450 lang=golang
 *
 * [450] Delete Node in a BST
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
func deleteNode(root *TreeNode, key int) *TreeNode {
	var helper func(node *TreeNode) *TreeNode
	helper = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		if node.Val < key {
			node.Right = helper(node.Right)
			return node
		}
		if node.Val > key {
			node.Left = helper(node.Left)
			return node
		}
		if node.Val == key {
			if node.Left != nil && node.Right != nil {
				right := getRightNodeDeepestLeft(node.Right)
				right.Left = node.Left
				node.Left = nil
				return node.Right
			}
			if node.Left != nil && node.Right == nil {
				return node.Left
			}
			if node.Left == nil && node.Right != nil {
				return node.Right
			}
			return nil
		}
		node.Left = helper(node.Left)
		node.Right = helper(node.Right)
		return node
	}
	return helper(root)
}
func getRightNodeDeepestLeft(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

// @lc code=end
