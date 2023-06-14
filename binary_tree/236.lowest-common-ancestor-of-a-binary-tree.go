package main

/*
 * @lc app=leetcode id=236 lang=golang
 *
 * [236] Lowest Common Ancestor of a Binary Tree
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
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	ancestor := &TreeNode{}
	var helper func(node *TreeNode) bool
	helper = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		leftFound := helper(node.Left)
		rightFound := helper(node.Right)

		if leftFound || rightFound {
			if (node == p || node == q) || (leftFound == rightFound) {
				ancestor = node
			}
			return true
		}
		if node == p || node == q {
			return true
		}
		return false
	}
	helper(root)

	return ancestor
}

// @lc code=end
