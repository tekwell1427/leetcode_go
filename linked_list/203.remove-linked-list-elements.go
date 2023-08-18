package main

/*
 * @lc app=leetcode id=203 lang=golang
 *
 * [203] Remove Linked List Elements
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	pseudo := &ListNode{}
	pre := pseudo

	for head != nil {
		if head.Val != val {
			pre.Next = head
			pre = head
		}
		head = head.Next
		pre.Next = nil
	}
	return pseudo.Next
}

// @lc code=end
