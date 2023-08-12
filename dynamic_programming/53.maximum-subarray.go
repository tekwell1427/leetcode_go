package main

/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// @lc code=start
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums)+1)
	dp[0] = nums[0]
	maxSum := nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		dp[i] = max(dp[i-1]+num, num)
		maxSum = max(maxSum, dp[i])
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
