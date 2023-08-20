package main

/*
 * @lc app=leetcode id=142 lang=golang
 *
 * [142] Linked List Cycle II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			slow = head
			break
		}
	}

	for fast != nil && fast.Next != nil {
		if slow == fast {
			return slow
		}
		slow = slow.Next
		fast = fast.Next
	}
	return nil
}

// @lc code=end
