/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 */

// @lc code=start
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
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

