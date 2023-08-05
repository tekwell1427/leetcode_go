package main

/*
 * @lc app=leetcode id=416 lang=golang
 *
 * [416] Partition Equal Subset Sum
 */

// @lc code=start
func canPartition(nums []int) bool {
	numSum := 0
	for _, num := range nums {
		numSum += num
	}
	if numSum%2 != 0 {
		return false
	}
	mean := numSum / 2

	dp := make([]int, mean+1)

	for _, num := range nums {
		for curMean := mean; curMean >= 0; curMean-- {
			if curMean >= num {
				dp[curMean] = max(dp[curMean], dp[curMean-num]+num)
			}
		}
	}
	return dp[mean] == mean
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
