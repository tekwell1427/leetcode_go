package main

/*
 * @lc app=leetcode id=2130 lang=golang
 *
 * [2130] Maximum Twin Sum of a Linked List
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
func pairSum(head *ListNode) int {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	var preNode *ListNode
	curNode := slow

	for curNode != nil {
		nextNode := curNode.Next
		curNode.Next = preNode

		preNode = curNode
		curNode = nextNode
	}
	head2 := preNode
	maxSum := 0
	for head2 != nil {
		maxSum = max(maxSum, head.Val+head2.Val)
		head = head.Next
		head2 = head2.Next
	}
	return maxSum
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
