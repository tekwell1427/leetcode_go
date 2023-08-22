package main

/*
 * @lc app=leetcode id=239 lang=golang
 *
 * [239] Sliding Window Maximum
 */

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	deque := []int{}
	output := []int{}
	for i := 0; i < k; i++ {
		for len(deque) > 0 && deque[len(deque)-1] < nums[i] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, nums[i])
	}
	output = append(output, deque[0])
	for i := k; i < len(nums); i++ {
		if nums[i-k] == deque[0] {
			deque = deque[1:]
		}
		for len(deque) > 0 && deque[len(deque)-1] < nums[i] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, nums[i])
		output = append(output, deque[0])
	}
	return output
}

// @lc code=end
