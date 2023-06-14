package main

/*
 * @lc app=leetcode id=235 lang=golang
 *
 * [235] Lowest Common Ancestor of a Binary Search Tree
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
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val > q.Val {
		tmp := p
		p = q
		q = tmp
	}
	for {
		if root.Val < p.Val {
			root = root.Right
		} else if root.Val > q.Val {
			root = root.Left
		} else {
			break
		}
	}
	return root
}

// @lc code=end
