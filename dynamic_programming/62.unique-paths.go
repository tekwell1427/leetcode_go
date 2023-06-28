package main

/*
 * @lc app=leetcode id=62 lang=golang
 *
 * [62] Unique Paths
 */

// @lc code=start
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i, _ := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for colIdx := 0; colIdx < n; colIdx++ {
		dp[0][colIdx] = 1
	}
	for rowIdx := 1; rowIdx < m; rowIdx++ {
		for colIdx := 1; colIdx < n; colIdx++ {
			dp[rowIdx][colIdx] = dp[rowIdx-1][colIdx] + dp[rowIdx][colIdx-1]
		}
	}
	return dp[m-1][n-1]
}

// @lc code=end
