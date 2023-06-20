package main

import "fmt"

/*
 * @lc app=leetcode id=37 lang=golang
 *
 * [37] Sudoku Solver
 */

// @lc code=start
func solveSudoku(board [][]byte) {
	var helper func() bool
	helper = func() bool {
		for rowIdx := 0; rowIdx < 9; rowIdx++ {
			for colIdx := 0; colIdx < 9; colIdx++ {
				// until find empty '.'
				if board[rowIdx][colIdx] != '.' {
					continue
				}
				for i := '1'; i <= '9'; i++ {
					if isLegal(board, rowIdx, colIdx, byte(i)) {
						board[rowIdx][colIdx] = byte(i)
						if helper() {
							return true
						}
						board[rowIdx][colIdx] = '.'
					}
				}
				// all num cannot fit, means error on pre-step
				return false
			}
		}
		// there is no empty '.' to assign num, means finishing
		return true
	}
	helper()
}
func isLegal(b [][]byte, row int, col int, num byte) bool {
	for rowIdx := 0; rowIdx < 9; rowIdx++ {
		if b[rowIdx][col] == num {
			return false
		}
	}
	for colIdx := 0; colIdx < 9; colIdx++ {
		if b[row][colIdx] == num {
			return false
		}
	}
	rowIdx := row / 3 * 3
	colIdx := col / 3 * 3
	fmt.Println(rowIdx, colIdx, "^^")
	for i := rowIdx; i < rowIdx+3; i++ {
		for j := colIdx; j < colIdx+3; j++ {
			if b[i][j] == num {
				return false
			}
		}
	}
	return true
}

// @lc code=end
