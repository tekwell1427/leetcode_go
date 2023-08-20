package main

/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow, fast := head, head

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	dummy := &ListNode{Next: head}
	pre := dummy
	for fast != nil {
		fast = fast.Next
		pre = slow
		slow = slow.Next
	}
	pre.Next = slow.Next

	return dummy.Next
}

// @lc code=end
