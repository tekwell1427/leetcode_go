package main

/*
 * @lc app=leetcode id=122 lang=golang
 *
 * [122] Best Time to Buy and Sell Stock II
 */

// @lc code=start
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	sum := 0
	for i := 1; i < len(prices); i++ {
		profit := prices[i] - prices[i-1]
		if profit > 0 {
			sum += profit
		}
	}
	return sum
}

// @lc code=end
