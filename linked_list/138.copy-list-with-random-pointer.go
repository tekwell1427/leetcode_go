package main

/*
 * @lc app=leetcode id=138 lang=golang
 *
 * [138] Copy List with Random Pointer
 */
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	old := head

	for old != nil {
		new := &Node{Val: old.Val}
		nextOld := old.Next
		old.Next = new
		new.Next = nextOld
		old = nextOld
	}
	old = head
	for old != nil {

		if old.Random != nil {
			old.Next.Random = old.Random.Next
		}

		old = old.Next.Next
	}

	newHead := head.Next
	old = head
	new := newHead
	for old != nil {
		old.Next = new.Next
		old = old.Next
		// if old is end now
		if old != nil {
			new.Next = new.Next.Next
			new = new.Next
		}
	}
	return newHead

}

// @lc code=end
