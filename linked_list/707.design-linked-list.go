package main

/*
 * @lc app=leetcode id=707 lang=golang
 *
 * [707] Design Linked List
 */

// @lc code=start
type Node struct {
	pre  *Node
	next *Node
	val  int
}
type MyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func Constructor() MyLinkedList {
	pseudo := &Node{}
	pseudo.next = pseudo
	pseudo.pre = pseudo
	return MyLinkedList{
		head:   pseudo,
		tail:   pseudo,
		length: 0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index >= this.length {
		return -1
	}
	cur := this.head
	for i := 0; i <= index; i++ {
		cur = cur.next
	}
	return cur.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	new := &Node{val: val, pre: this.head, next: this.head.next}
	this.head.next.pre = new
	this.head.next = new

	this.length++
}

func (this *MyLinkedList) AddAtTail(val int) {
	new := &Node{val: val, pre: this.tail.pre, next: this.tail}
	this.tail.pre.next = new
	this.tail.pre = new

	this.length++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		this.AddAtHead(val)
	} else if index == this.length {
		this.AddAtTail(val)
	} else if index < this.length {
		cur := this.head
		for i := 0; i <= index; i++ {
			cur = cur.next
		}
		new := &Node{val: val, pre: cur.pre, next: cur}
		cur.pre.next = new
		cur.pre = new

		this.length++
	}

}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < this.length {
		cur := this.head
		for i := 0; i <= index; i++ {
			cur = cur.next
		}
		cur.pre.next = cur.next
		cur.next.pre = cur.pre
		this.length--
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
// @lc code=end
