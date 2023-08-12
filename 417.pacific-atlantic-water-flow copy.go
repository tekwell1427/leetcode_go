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
