package main

/*
 * @lc app=leetcode id=328 lang=golang
 *
 * [328] Odd Even Linked List
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
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	evenHead := head.Next
	curEvenNode := head.Next
	curOddNode := head

	for curEvenNode != nil && curEvenNode.Next != nil {
		nextEvenNode := curEvenNode.Next.Next
		nextOddNode := curOddNode.Next.Next

		curEvenNode.Next = nextEvenNode
		curEvenNode = nextEvenNode

		curOddNode.Next = nextOddNode
		curOddNode = nextOddNode
	}
	curOddNode.Next = evenHead
	return head
}

// @lc code=end
