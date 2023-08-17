package main

/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow, fast := head, head

	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	pseudo := &ListNode{}
	pre := pseudo
	for fast != nil {
		pre.Next = slow
		pre = slow
		slow = slow.Next
		fast = fast.Next
	}
	pre.Next = slow.Next
	return pseudo.Next
}

// @lc code=end
