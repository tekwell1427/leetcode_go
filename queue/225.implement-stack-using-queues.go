package main

/*
 * @lc app=leetcode id=225 lang=golang
 *
 * [225] Implement Stack using Queues
 */

// @lc code=start
type MyStack struct {
	queueIn  []int
	queueOut []int
}

func Constructor() MyStack {
	return MyStack{
		queueIn:  []int{},
		queueOut: []int{},
	}
}

func (this *MyStack) Push(x int) {
	this.queueIn = append(this.queueIn, x)
	for _, element := range this.queueOut {
		this.queueIn = append(this.queueIn, element)
	}
	this.queueOut = this.queueIn
	this.queueIn = []int{}
}

func (this *MyStack) Pop() int {
	output := this.queueOut[0]
	this.queueOut = this.queueOut[1:]
	return output
}

func (this *MyStack) Top() int {
	return this.queueOut[0]
}

func (this *MyStack) Empty() bool {
	return len(this.queueOut) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end
