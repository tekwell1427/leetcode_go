package main

/*
 * @lc app=leetcode id=643 lang=golang
 *
 * [643] Maximum Average Subarray I
 */

// @lc code=start
func findMaxAverage(nums []int, k int) float64 {
	headIdx, curSum := 0, 0

	for i := 0; i < k; i++ {
		curSum += nums[i]
	}
	maxSum := curSum

	for i := k; i < len(nums); i++ {
		curSum -= nums[headIdx]
		curSum += nums[i]
		headIdx++

		maxSum = max(maxSum, curSum)
	}
	return float64(maxSum) / float64(k)
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
