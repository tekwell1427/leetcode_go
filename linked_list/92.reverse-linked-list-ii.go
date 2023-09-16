package main

/*
 * @lc app=leetcode id=92 lang=golang
 *
 * [92] Reverse Linked List II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	cur := head
	for i := 0; i < left-1; i++ {
		pre = cur
		cur = cur.Next
	}
	// pre is node which will connect with reverse linked list
	for i := 0; i < right-left; i++ {
		moveToHead := cur.Next
		newNext := cur.Next.Next
		cur.Next = newNext

		moveToHead.Next = pre.Next
		pre.Next = moveToHead
	}
	// cur is tail, not including reverse node

	return dummy.Next
}

// @lc code=end
