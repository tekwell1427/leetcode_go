package main

/*
 * @lc app=leetcode id=63 lang=golang
 *
 * [63] Unique Paths II
 */

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	maxRow, maxCol := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, len(obstacleGrid))
	for rowIdx := 0; rowIdx < maxRow; rowIdx++ {
		dp[rowIdx] = make([]int, len(obstacleGrid[0]))
		if rowIdx > 0 && obstacleGrid[rowIdx-1][0] == 1 {
			obstacleGrid[rowIdx][0] = 1
			continue
		}
		if obstacleGrid[rowIdx][0] == 0 {
			dp[rowIdx][0] = 1
		}

	}
	for colIdx := 0; colIdx < maxCol; colIdx++ {
		if colIdx > 0 && obstacleGrid[0][colIdx-1] == 1 {
			break
		}
		if obstacleGrid[0][colIdx] == 0 {
			dp[0][colIdx] = 1
		}
	}
	for rowIdx := 1; rowIdx < maxRow; rowIdx++ {
		for colIdx := 1; colIdx < maxCol; colIdx++ {
			if obstacleGrid[rowIdx][colIdx] == 1 {
				continue
			}
			dp[rowIdx][colIdx] = dp[rowIdx][colIdx-1] + dp[rowIdx-1][colIdx]
		}
	}
	return dp[maxRow-1][maxCol-1]
}

// @lc code=end
