package main

/*
 * @lc app=leetcode id=213 lang=golang
 *
 * [213] House Robber II
 */

// @lc code=start
func rob(nums []int) int {
	if len(nums) <= 1 {
		return nums[0]
	}
	max1 := getMaxMoney(nums[:len(nums)-1]) // without TAIL num
	max2 := getMaxMoney(nums[1:len(nums)])  // without HEAD num
	return max(max1, max2)
}
func getMaxMoney(nums []int) int {
	if len(nums) <= 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		num := nums[i]
		dp[i] = max(dp[i-1], dp[i-2]+num)
	}
	return dp[len(dp)-1]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
