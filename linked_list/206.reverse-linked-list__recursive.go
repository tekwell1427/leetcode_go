package main

/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	var helper func(pre, cur *ListNode) *ListNode
	helper = func(pre, cur *ListNode) *ListNode {
		if cur == nil {
			return pre
		}
		next := cur.Next
		cur.Next = pre
		return helper(cur, next)
	}
	return helper(pre, head)
}

// @lc code=end
