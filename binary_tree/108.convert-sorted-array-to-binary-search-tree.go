package main

/*
 * @lc app=leetcode id=108 lang=golang
 *
 * [108] Convert Sorted Array to Binary Search Tree
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
func sortedArrayToBST(nums []int) *TreeNode {
	var helper func(node *TreeNode, subNums []int) *TreeNode
	helper = func(node *TreeNode, subNums []int) *TreeNode {
		if len(subNums) == 0 {
			return nil
		}
		midIdx := len(subNums) / 2
		newNode := &TreeNode{Val: subNums[midIdx]}
		newNode.Left = helper(newNode, subNums[:midIdx])
		newNode.Right = helper(newNode, subNums[midIdx+1:])
		return newNode
	}
	return helper(new(TreeNode), nums)
}

// func getMidIdx(subNums []int) {
// 	len(subNums) / 2
// }

// @lc code=end
