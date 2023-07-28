package main

/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	rowMax, colMax := len(grid), len(grid[0])

	var helper func(rowIdx, colIdx int) bool
	helper = func(rowIdx, colIdx int) bool {
		if rowIdx < 0 || rowIdx == rowMax || colIdx < 0 || colIdx == colMax {
			return false
		}

		if grid[rowIdx][colIdx] == '1' {
			queue := [][2]int{}
			queue = append(queue, [2]int{rowIdx, colIdx})

			for len(queue) > 0 {
				point := queue[0]
				queue = queue[1:]
				rowIdx, colIdx := point[0], point[1]

				if rowIdx < 0 || rowIdx == rowMax || colIdx < 0 || colIdx == colMax {
					continue
				}
				if grid[rowIdx][colIdx] == '2' {
					continue
				}
				if grid[rowIdx][colIdx] == '0' {
					continue
				}
				grid[rowIdx][colIdx] = '2'
				queue = append(queue, [2]int{rowIdx - 1, colIdx})
				queue = append(queue, [2]int{rowIdx + 1, colIdx})
				queue = append(queue, [2]int{rowIdx, colIdx - 1})
				queue = append(queue, [2]int{rowIdx, colIdx + 1})
			}
			return true
		}
		return false
	}

	isLandCount := 0
	for rowIdx := 0; rowIdx < rowMax; rowIdx++ {
		for colIdx := 0; colIdx < colMax; colIdx++ {
			if helper(rowIdx, colIdx) {
				isLandCount++
			}
		}
	}
	return isLandCount
}

// @lc code=end
