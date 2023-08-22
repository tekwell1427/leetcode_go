package main

import "container/heap"

/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 */

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	numCountMap := make(map[int]int)
	for _, num := range nums {
		numCountMap[num]++
	}

	my := &myHeap{
		nums: [][2]int{},
	}
	heap.Init(my)
	for num, count := range numCountMap {
		heap.Push(my, [2]int{num, count})
		if my.Len() > k {
			heap.Pop(my)
		}
	}

	output := []int{}
	for i := 0; i < k; i++ {
		output = append(output, heap.Pop(my).(int))
	}

	// my.nums only has top k:
	return output
}

type myHeap struct {
	nums [][2]int
}

func (this *myHeap) Push(x any) {
	y := x.([2]int)
	this.nums = append(this.nums, y)
}

// remove and return element Len() - 1.
func (this *myHeap) Pop() any {
	output := this.nums[len(this.nums)-1]
	this.nums = this.nums[:len(this.nums)-1]
	return output[0]
}

func (this *myHeap) Len() int {
	return len(this.nums)
}
func (this *myHeap) Less(i, j int) bool {
	return this.nums[i][1] < this.nums[j][1]
}
func (this *myHeap) Swap(i, j int) {
	this.nums[i], this.nums[j] = this.nums[j], this.nums[i]
}

// @lc code=end
