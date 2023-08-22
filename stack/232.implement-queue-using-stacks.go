package main

/*
 * @lc app=leetcode id=232 lang=golang
 *
 * [232] Implement Queue using Stacks
 */

// @lc code=start
type MyQueue struct {
	stackIn  []int
	stackOut []int
}

func Constructor() MyQueue {
	return MyQueue{
		stackIn:  []int{},
		stackOut: []int{},
	}
}

func (this *MyQueue) Push(x int) {
	this.stackIn = append(this.stackIn, x)
}

func (this *MyQueue) Pop() int {
	if len(this.stackOut) == 0 {
		for idx := len(this.stackIn) - 1; idx >= 0; idx-- {
			element := this.stackIn[idx]
			this.stackOut = append(this.stackOut, element)
		}
		this.stackIn = []int{}
	}
	output := this.stackOut[len(this.stackOut)-1]
	this.stackOut = this.stackOut[:len(this.stackOut)-1]
	return output
}

func (this *MyQueue) Peek() int {
	if len(this.stackOut) == 0 {
		for idx := len(this.stackIn) - 1; idx >= 0; idx-- {
			element := this.stackIn[idx]
			this.stackOut = append(this.stackOut, element)
		}
		this.stackIn = []int{}
	}
	return this.stackOut[len(this.stackOut)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.stackIn) == 0 && len(this.stackOut) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end
