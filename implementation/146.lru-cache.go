package main

/*
 * @lc app=leetcode id=146 lang=golang
 *
 * [146] LRU Cache
 */

// @lc code=start
type DoublyLinkedNode struct {
	pre  *DoublyLinkedNode
	next *DoublyLinkedNode
	val  int
	key  int
}
type LRUCache struct {
	record   map[int]*DoublyLinkedNode
	head     *DoublyLinkedNode
	tail     *DoublyLinkedNode
	capacity int
	length   int
}

func Constructor(capacity int) LRUCache {
	dummy := &DoublyLinkedNode{}
	dummy.pre = dummy
	dummy.next = dummy
	return LRUCache{
		record:   make(map[int]*DoublyLinkedNode),
		head:     dummy,
		tail:     dummy,
		capacity: capacity,
		length:   0,
	}
}

func (this *LRUCache) Get(key int) int {
	output := -1
	if node, exits := this.record[key]; exits {
		output = node.val
		this.DeleteNode(node)
		this.AddNodeTohead(node)
	}
	return output
}

func (this *LRUCache) Put(key int, value int) {
	if node, exits := this.record[key]; exits {
		node.val = value
		this.DeleteNode(node)
		this.AddNodeTohead(node)
	} else {
		node := &DoublyLinkedNode{key: key, val: value}
		this.record[key] = node
		this.AddNodeTohead(node)
		this.length++
	}
	if this.length > this.capacity {
		delete(this.record, this.tail.pre.key)
		this.DeleteNode(this.tail.pre)
		this.length--
	}
}
func (this *LRUCache) DeleteNode(node *DoublyLinkedNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
	node.next, node.pre = nil, nil
}
func (this *LRUCache) AddNodeTohead(node *DoublyLinkedNode) {
	this.head.next.pre = node
	node.next = this.head.next
	this.head.next = node
	node.pre = this.head
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
