package main

/*
 * @lc app=leetcode id=64 lang=golang
 *
 * [64] Minimum Path Sum
 */

// @lc code=start
func minPathSum(grid [][]int) int {
	rowMax, colMax := len(grid), len(grid[0])
	dp := make([][]int, len(grid))
	for rowIdx := range dp {
		dp[rowIdx] = make([]int, len(grid[0]))
	}
	dp[0][0] = grid[0][0]
	for rowIdx := 1; rowIdx < rowMax; rowIdx++ {
		dp[rowIdx][0] = dp[rowIdx-1][0] + grid[rowIdx][0]
	}

	for colIdx := 1; colIdx < colMax; colIdx++ {
		dp[0][colIdx] += dp[0][colIdx-1] + grid[0][colIdx]
	}

	for rowIdx := 1; rowIdx < rowMax; rowIdx++ {
		for colIdx := 1; colIdx < colMax; colIdx++ {
			dp[rowIdx][colIdx] = min(dp[rowIdx-1][colIdx], dp[rowIdx][colIdx-1]) + grid[rowIdx][colIdx]
		}
	}
	return dp[rowMax-1][colMax-1]
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end
