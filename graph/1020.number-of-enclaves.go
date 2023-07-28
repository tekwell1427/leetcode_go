package main

/*
 * @lc app=leetcode id=1020 lang=golang
 *
 * [1020] Number of Enclaves
 */

// @lc code=start
func numEnclaves(grid [][]int) int {
	rowMax, colMax := len(grid), len(grid[0])
	var helper func(rowIdx, colIdx int) int
	helper = func(rowIdx, colIdx int) int {
		if rowIdx < 0 || rowIdx == rowMax || colIdx < 0 || colIdx == colMax {
			// return boundary signal
			return -1
		}
		if grid[rowIdx][colIdx] == 1 {
			grid[rowIdx][colIdx] = 2

			top := helper(rowIdx-1, colIdx)
			buttom := helper(rowIdx+1, colIdx)
			left := helper(rowIdx, colIdx-1)
			right := helper(rowIdx, colIdx+1)
			if top == -1 || buttom == -1 || left == -1 || right == -1 {
				// return boundary signal
				return -1
			}
			return top + buttom + left + right + 1
		}
		return 0
	}

	enclaves := 0
	for rowIdx := 0; rowIdx < rowMax; rowIdx++ {
		for colIdx := 0; colIdx < colMax; colIdx++ {
			area := helper(rowIdx, colIdx)
			if area != -1 {
				enclaves += area
			}
		}
	}
	return enclaves
}

// @lc code=end
