package main

import (
	"math"
)

/*
 * @lc app=leetcode id=417 lang=golang
 *
 * [417] Pacific Atlantic Water Flow
 */
// -5: cannot flow anywhere
// -1: pacific ocean
// -2: atlantic ocean
// -3: both
// @lc code=start
func pacificAtlantic(heights [][]int) [][]int {
	output := [][]int{}
	rowMax, colMax := len(heights), len(heights[0])
	visited := make([][]int, rowMax)
	for i := 0; i < rowMax; i++ {
		visited[i] = make([]int, colMax)
	}

	var helper func(rowIdx, colIdx int, preNum int) int
	helper = func(rowIdx, colIdx int, preNum int) int {
		if rowIdx < 0 || colIdx < 0 {
			return -1
		}
		if rowIdx == rowMax || colIdx == colMax {
			return -2
		}
		if visited[rowIdx][colIdx] < 0 {
			return visited[rowIdx][colIdx]
		}
		if heights[rowIdx][colIdx] == -4 {
			return -5
		}

		if heights[rowIdx][colIdx] <= preNum {
			curNum := heights[rowIdx][colIdx]
			heights[rowIdx][colIdx] = -4

			top := helper(rowIdx-1, colIdx, curNum)
			buttom := helper(rowIdx+1, colIdx, curNum)
			left := helper(rowIdx, colIdx-1, curNum)
			right := helper(rowIdx, colIdx+1, curNum)

			heights[rowIdx][colIdx] = curNum

			result := []int{top, buttom, left, right}
			pacific, atlantic := false, false
			for _, i := range result {
				switch i {
				case -1:
					pacific = true
				case -2:
					atlantic = true
				case -3:
					pacific, atlantic = true, true
				}
			}
			var record int
			if pacific && atlantic {
				record = -3
			} else if pacific {
				record = -1
			} else if atlantic {
				record = -2
			} else {
				record = -5
			}
			visited[rowIdx][colIdx] = record
			return record
		}
		return -5
	}
	for rowIdx := 0; rowIdx < rowMax; rowIdx++ {
		for colIdx := 0; colIdx < colMax; colIdx++ {
			if helper(rowIdx, colIdx, math.MaxInt32) == -3 {
				output = append(output, []int{rowIdx, colIdx})
			}
		}
	}
	return output
}

// @lc code=end
