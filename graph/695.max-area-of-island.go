package main

/*
 * @lc app=leetcode id=695 lang=golang
 *
 * [695] Max Area of Island
 */

// @lc code=start
func maxAreaOfIsland(grid [][]int) int {
	rowMax, colMax := len(grid), len(grid[0])
	var helper func(rowIdx, colIdx int) int
	helper = func(rowIdx, colIdx int) int {
		if rowIdx < 0 || rowIdx == rowMax || colIdx < 0 || colIdx == colMax {
			return 0
		}
		if grid[rowIdx][colIdx] == 1 {
			grid[rowIdx][colIdx] = 2
			return 1 +
				helper(rowIdx-1, colIdx) +
				helper(rowIdx+1, colIdx) +
				helper(rowIdx, colIdx-1) +
				helper(rowIdx, colIdx+1)
		}
		return 0
	}

	maxArea := 0
	for rowIdx := 0; rowIdx < rowMax; rowIdx++ {
		for colIdx := 0; colIdx < colMax; colIdx++ {
			maxArea = max(maxArea, helper(rowIdx, colIdx))
		}
	}

	return maxArea
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
