package main

/*
 * @lc app=leetcode id=24 lang=golang
 *
 * [24] Swap Nodes in Pairs
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy

	for head != nil && head.Next != nil {
		next := head.Next // 2
		pre.Next = next   // 0 2
		newHead := next.Next
		next.Next = head
		pre = head
		head = newHead
		pre.Next = head
	}

	return dummy.Next
}

// @lc code=end
