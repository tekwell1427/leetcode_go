package main

/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
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
func reverseList(head *ListNode) *ListNode {
	var dummy *ListNode
	preNode := dummy
	curNode := head
	for curNode != nil {
		nextNode := curNode.Next
		curNode.Next = preNode

		preNode = curNode
		curNode = nextNode
	}
	return preNode
}

// @lc code=end
