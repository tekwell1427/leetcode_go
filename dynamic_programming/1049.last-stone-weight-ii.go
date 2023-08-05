package main

/*
 * @lc app=leetcode id=1049 lang=golang
 *
 * [1049] Last Stone Weight II
 */

// @lc code=start
func lastStoneWeightII(stones []int) int {
	stoneSum := 0
	for _, stone := range stones {
		stoneSum += stone
	}
	mean := stoneSum / 2

	dp := make([]int, mean+1)

	for _, stone := range stones {
		for curMean := mean; curMean >= 0; curMean-- {
			if curMean >= stone {
				dp[curMean] = max(dp[curMean], dp[curMean-stone]+stone)
			}
		}
	}
	return stoneSum - 2*dp[mean]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
