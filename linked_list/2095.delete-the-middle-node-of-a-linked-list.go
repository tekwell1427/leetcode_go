package main

/*
 * @lc app=leetcode id=2095 lang=golang
 *
 * [2095] Delete the Middle Node of a Linked List
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
func deleteMiddle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head, head
	var preNode *ListNode
	for fast != nil && fast.Next != nil {
		preNode = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	preNode.Next = slow.Next
	return head
}

// @lc code=end
