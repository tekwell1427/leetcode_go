package main

/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// @lc code=start
func maxSubArray(nums []int) int {
	curSum := 0
	maxSum := nums[0]
	for _, num := range nums {
		curSum += num
		maxSum = max(maxSum, curSum)

		if curSum < 0 {
			curSum = 0
		}
	}
	return maxSum
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
