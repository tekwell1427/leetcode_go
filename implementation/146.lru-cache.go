package main

/*
 * @lc app=leetcode id=146 lang=golang
 *
 * [146] LRU Cache
 */

// @lc code=start
type LRUCache struct {
	head, tail *DLinkedNode
	cache      map[int]*DLinkedNode
	capacity   int
	size       int
}

type DLinkedNode struct {
	pre, next *DLinkedNode
	key, val  int
}

func Constructor(capacity int) LRUCache {
	head := &DLinkedNode{key: 0, val: 0}
	tail := &DLinkedNode{key: 0, val: 0}
	head.next = tail
	tail.pre = head
	return LRUCache{
		head:     head,
		tail:     tail,
		capacity: capacity,
		size:     0,
		cache:    map[int]*DLinkedNode{},
	}
}

func (this *LRUCache) Get(key int) int {
	if node, exists := this.cache[key]; exists {
		this.moveToHead(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, exists := this.cache[key]; exists {
		node.val = value
		this.moveToHead(node)
	} else {
		newNode := &DLinkedNode{key: key, val: value}
		this.addToHead(newNode)
		this.cache[key] = newNode
		this.size++
		if this.size > this.capacity {
			delete(this.cache, this.tail.pre.key)
			this.deleteTailNode()
			this.size--
		}
	}
}
func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.deleteNode(node)
	this.addToHead(node)
}
func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.pre = this.head
	node.next = this.head.next
	this.head.next.pre = node
	this.head.next = node
}

func (this *LRUCache) deleteNode(node *DLinkedNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}
func (this *LRUCache) deleteTailNode() {
	this.deleteNode(this.tail.pre)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
