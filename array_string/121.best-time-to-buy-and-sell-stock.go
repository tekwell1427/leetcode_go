package main

import "math"

/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
func maxProfit(prices []int) int {
	preMin := math.MaxInt32
	output := 0

	for _, price := range prices {
		preMin = min(preMin, price)
		output = max(output, price-preMin)
	}
	return output
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
