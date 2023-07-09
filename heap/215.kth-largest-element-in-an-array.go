package main

import (
	"container/heap"
)

/*
 * @lc app=leetcode id=215 lang=golang
 *
 * [215] Kth Largest Element in an Array
 */

// @lc code=start
type myMinHeap struct {
	nums []int
}

func (m *myMinHeap) Push(x any) {
	m.nums = append(m.nums, x.(int))
}
func (m *myMinHeap) Pop() any {
	p := m.nums[len(m.nums)-1]
	m.nums = m.nums[:len(m.nums)-1]
	return p
}
func (m myMinHeap) Len() int {
	return len(m.nums)
}
func (m myMinHeap) Less(i, j int) bool {
	return m.nums[i] < m.nums[j]
}
func (m myMinHeap) Swap(i, j int) {
	m.nums[i], m.nums[j] = m.nums[j], m.nums[i]
}

func findKthLargest(nums []int, k int) int {
	myheap := &myMinHeap{nums: []int{}}
	heap.Init(myheap)

	for _, num := range nums {
		heap.Push(myheap, num)
		if myheap.Len() > k {
			heap.Pop(myheap)
		}
	}
	return heap.Pop(myheap).(int)
}

// @lc code=end
