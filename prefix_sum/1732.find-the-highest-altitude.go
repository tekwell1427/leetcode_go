package main

/*
 * @lc app=leetcode id=1732 lang=golang
 *
 * [1732] Find the Highest Altitude
 */

// @lc code=start
func largestAltitude(gain []int) int {
	maxA := 0
	sum := 0
	for _, g := range gain {
		sum += g
		maxA = max(maxA, sum)
	}
	return maxA
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
