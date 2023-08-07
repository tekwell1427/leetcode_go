package main

/*
 * @lc app=leetcode id=518 lang=golang
 *
 * [518] Coin Change II
 */

// @lc code=start
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	// set one is meaninigful, when curAmount == curCoin, there has one number of arrangement
	dp[0] = 1

	for _, curCoin := range coins {
		for curAmount := curCoin; curAmount <= amount; curAmount++ {
			dp[curAmount] = dp[curAmount] + dp[curAmount-curCoin]
		}
	}

	return dp[amount]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
