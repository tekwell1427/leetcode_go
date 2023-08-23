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
	queue := []*TreeNode{root.Left, root.Right}

	for len(queue) > 0 {
		left, right := queue[0], queue[1]
		queue = queue[2:]

		if left == nil && right == nil {
			continue
		}
		if (left == nil && right != nil) || (left != nil && right == nil) {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		queue = append(queue, left.Left)
		queue = append(queue, right.Right)
		queue = append(queue, left.Right)
		queue = append(queue, right.Left)
	}
	return true
}

// @lc code=end
