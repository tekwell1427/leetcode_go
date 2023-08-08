package main

/*
 * @lc app=leetcode id=377 lang=golang
 *
 * [377] Combination Sum IV
 */

// @lc code=start
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	for curTarget := 0; curTarget <= target; curTarget++ {
		for _, num := range nums {
			if curTarget >= num {
				dp[curTarget] += dp[curTarget-num]
			}
		}
	}

	return dp[target]
}

// @lc code=end
