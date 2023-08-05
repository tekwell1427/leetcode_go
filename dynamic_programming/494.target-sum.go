package main

/*
 * @lc app=leetcode id=494 lang=golang
 *
 * [494] Target Sum
 */

// @lc code=start
func findTargetSumWays(nums []int, target int) int {
	numSum := 0
	for _, num := range nums {
		numSum += num
	}
	if numSum < abs(target) {
		return 0
	}
	if (target+numSum)%2 == 1 {
		return 0
	}
	left := (target + numSum) / 2

	dp := make([]int, left+1)

	// dp[nums[0]] = 1
	dp[0] = 1

	for _, num := range nums {
		for curLeft := left; curLeft >= 0; curLeft-- {
			if curLeft >= num {
				// dp[curLeft] is previous number of arrange set
				dp[curLeft] = dp[curLeft] + dp[curLeft-num]
			}
		}
	}
	return dp[left]
}
func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

// @lc code=end
