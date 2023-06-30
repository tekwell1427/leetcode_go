package main

/*
 * @lc app=leetcode id=1049 lang=golang
 *
 * [1049] Last Stone Weight II
 */

// @lc code=start
func lastStoneWeightII(stones []int) int {
	var sum int
	for _, stone := range stones {
		sum += stone
	}
	mean := sum / 2
	dp := make([]int, mean+1)

	for objectIdx := 0; objectIdx < len(stones); objectIdx++ {
		stone := stones[objectIdx]
		for bagSize := mean; bagSize >= stone; bagSize-- {
			if dp[bagSize-stone]+stone <= mean {
				dp[bagSize] = max(dp[bagSize], dp[bagSize-stone]+stone)
			}
		}
	}
	return sum - 2*dp[len(dp)-1]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
