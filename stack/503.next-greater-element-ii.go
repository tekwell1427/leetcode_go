package main

/*
 * @lc app=leetcode id=503 lang=golang
 *
 * [503] Next Greater Element II
 */

// @lc code=start
func nextGreaterElements(nums []int) []int {
	nLen := len(nums)
	output := make([]int, nLen)
	stack := []int{}

	// set largest numbers are -1
	for i := 0; i < nLen; i++ {
		output[i] = -1
	}

	for i := 0; i < 2*nLen; i++ {
		num := nums[i%nLen]
		for len(stack) > 0 && nums[stack[len(stack)-1]] < num {
			preIdx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			output[preIdx] = num
		}
		stack = append(stack, i%nLen)
	}
	return output
}

// @lc code=end
