package main

/*
 * @lc app=leetcode id=106 lang=golang
 *
 * [106] Construct Binary Tree from Inorder and Postorder Traversal
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	var helper func(preInorder, prePostOrder []int) *TreeNode
	helper = func(preInorder, prePostOrder []int) *TreeNode {
		if len(preInorder) == 0 {
			return nil
		}
		if len(preInorder) == 1 {
			return &TreeNode{Val: preInorder[0]}
		}
		mid := prePostOrder[len(prePostOrder)-1]
		midNode := &TreeNode{Val: mid}
		leftNums := getIndex(mid, preInorder) // left, mid, right
		rightNums := len(preInorder) - 1 - leftNums

		midNode.Left = helper(preInorder[:leftNums], prePostOrder[:leftNums])
		midNode.Right = helper(preInorder[leftNums+1:], prePostOrder[leftNums:leftNums+rightNums])

		return midNode
	}
	return helper(inorder, postorder)
}
func getIndex(j int, a []int) int {
	var index int
	for i, element := range a {
		if j == element {
			index = i
			break
		}
	}
	return index
}

// @lc code=end
