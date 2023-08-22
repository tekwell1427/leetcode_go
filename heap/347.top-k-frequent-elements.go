package main

import (
	"container/heap"
)

/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 */

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	my := &myHeap{
		nums:    []int{},
		collect: make(map[int]int),
	}
	heap.Init(my)
	numCountMap := make(map[int]int)

	for _, num := range nums {
		numCountMap[num]++
	}
	for num, count := range numCountMap {
		heap.Push(my, [2]int{num, count})
		if my.Len() > k {
			heap.Pop(my)
		}
	}

	// output := []int{}
	// for i := 0; i < k; i++ {
	// 	output = append(output, heap.Pop(my).(int))
	// }

	// my.nums only has top k:
	return my.nums
}

type myHeap struct {
	nums    []int
	collect map[int]int
}

func (this *myHeap) Push(x any) {
	y := x.([2]int)
	if _, exists := this.collect[y[0]]; !exists {
		this.nums = append(this.nums, y[0])
	}
	this.collect[y[0]] = y[1]
}

// remove and return element Len() - 1.
func (this *myHeap) Pop() any {
	output := this.nums[len(this.nums)-1]
	this.nums = this.nums[:len(this.nums)-1]
	delete(this.collect, output)
	return output
}

func (this *myHeap) Len() int {
	return len(this.nums)
}
func (this *myHeap) Less(i, j int) bool {
	return this.collect[this.nums[i]] < this.collect[this.nums[j]]
}
func (this *myHeap) Swap(i, j int) {
	this.nums[i], this.nums[j] = this.nums[j], this.nums[i]
}

// @lc code=end
