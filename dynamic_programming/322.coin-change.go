package main

import (
	"math"
)

/*
 * @lc app=leetcode id=322 lang=golang
 *
 * [322] Coin Change
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp2 := make([]int, amount+1)

	for i := 1; i < amount+1; i++ {
		dp[i] = math.MaxUint32
	}
	dp[0] = 0

	for _, curCoin := range coins {
		for curAmount := curCoin; curAmount <= amount; curAmount++ {
			dp[curAmount] = min(dp[curAmount], dp[curAmount-curCoin]+1)
			dp2[curAmount] = max(dp2[curAmount], dp2[curAmount-curCoin]+curCoin)
		}
	}
	if dp2[amount] != amount {
		return -1
	}
	return dp[amount]
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
