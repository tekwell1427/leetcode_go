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

	for i := 1; i < amount+1; i++ {
		dp[i] = math.MaxUint32
	}
	dp[0] = 0

	for _, curCoin := range coins {
		for curAmount := curCoin; curAmount <= amount; curAmount++ {
			dp[curAmount] = min(dp[curAmount], dp[curAmount-curCoin]+1)
		}
	}
	// if stone 2, and bagSize is 3， there is no any possible to acheive 3(bagsize)
	if dp[amount] == math.MaxUint32 {
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

// @lc code=end
