package main

import "container/heap"

/*
 * @lc app=leetcode id=215 lang=golang
 *
 * [215] Kth Largest Element in an Array
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {
	myHeap := &H{}
	heap.Init(myHeap)
	for _, num := range nums {
		heap.Push(myHeap, num)
		if len(myHeap.nums) > k {
			heap.Pop(myHeap)
		}
	}
	return heap.Pop(myHeap).(int)
}

type H struct {
	nums []int
}

func (h *H) Len() int {
	return len(h.nums)
}
func (h *H) Less(i, j int) bool {
	return h.nums[i] < h.nums[j]
}
func (h *H) Swap(i, j int) {
	h.nums[i], h.nums[j] = h.nums[j], h.nums[i]
}
func (h *H) Push(x any) {
	h.nums = append(h.nums, x.(int))
}
func (h *H) Pop() any {
	output := h.nums[len(h.nums)-1]
	h.nums = h.nums[:len(h.nums)-1]
	return output
}

// @lc code=end
