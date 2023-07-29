package main

/*
 * @lc app=leetcode id=130 lang=golang
 *
 * [130] Surrounded Regions
 */

// @lc code=start
func solve(board [][]byte) {
	rowMax, colMax := len(board), len(board[0])
	var helper func(rowIdx, colIdx int)
	helper = func(rowIdx, colIdx int) {
		if rowIdx < 0 || rowIdx == rowMax || colIdx < 0 || colIdx == colMax {
			// touch edge
			return
		}
		if board[rowIdx][colIdx] == 'X' {
			return
		}
		if board[rowIdx][colIdx] == 'O' {
			board[rowIdx][colIdx] = 'Y'
			helper(rowIdx-1, colIdx)
			helper(rowIdx+1, colIdx)
			helper(rowIdx, colIdx-1)
			helper(rowIdx, colIdx+1)
		}
	}
	// left, right
	for rowIdx := 0; rowIdx < rowMax; rowIdx++ {
		helper(rowIdx, 0)
		helper(rowIdx, colMax-1)
	}
	// top, buttom
	for colIdx := 0; colIdx < colMax; colIdx++ {
		helper(0, colIdx)
		helper(rowMax-1, colIdx)
	}

	for rowIdx := 0; rowIdx < rowMax; rowIdx++ {
		for colIdx := 0; colIdx < colMax; colIdx++ {
			if board[rowIdx][colIdx] == 'Y' {
				board[rowIdx][colIdx] = 'O'
			} else if board[rowIdx][colIdx] == 'O' {
				board[rowIdx][colIdx] = 'X'
			}
		}
	}
}

// @lc code=end
