package main

import (
	"strings"
)

/*
 * @lc app=leetcode id=51 lang=golang
 *
 * [51] N-Queens
 */

// @lc code=start
func solveNQueens(n int) [][]string {
	output := [][]string{}
	var helper func(preRowIdx int, subOutput []string)
	helper = func(preRowIdx int, subOutput []string) {
		if preRowIdx == n-1 {
			tmp := make([]string, len(subOutput))
			copy(tmp, subOutput)
			output = append(output, tmp)
			return
		}

		preRowIdx++
		preColIdx := 0
		for preColIdx < n {
			if !isConflict(subOutput, preColIdx, n) {
				s := strings.Repeat(".", preColIdx) + "Q" + strings.Repeat(".", n-preColIdx-1)
				subOutput = append(subOutput, s)
				helper(preRowIdx, subOutput)
				subOutput = subOutput[:len(subOutput)-1]
			}
			preColIdx++
		}
	}
	helper(-1, []string{})
	return output
}
func isConflict(output []string, col int, dimension int) bool {
	left, right := col-1, col+1
	for rowIdx := len(output) - 1; rowIdx >= 0; rowIdx-- {
		if output[rowIdx][col] == 'Q' {
			return true
		}
		if left >= 0 && output[rowIdx][left] == 'Q' {
			return true
		}
		if right < dimension && output[rowIdx][right] == 'Q' {
			return true
		}
		left--
		right++
	}

	return false
}

// @lc code=end
