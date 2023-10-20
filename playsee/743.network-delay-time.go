package main

import (
	"math"
)

/*
 * @lc app=leetcode id=743 lang=golang
 *
 * [743] Network Delay Time
 */

// @lc code=start
func networkDelayTime(times [][]int, n int, k int) int {
	// build start->arrive matrix
	sourceTarget := make([][]int, n)
	for i := 0; i < n; i++ {
		sourceTarget[i] = make([]int, n)
		for j := 0; j < n; j++ {
			sourceTarget[i][j] = math.MaxInt8
		}
	}
	for _, time := range times {
		srcNodeIdx, trgNodeIdx, t := time[0]-1, time[1]-1, time[2]
		sourceTarget[srcNodeIdx][trgNodeIdx] = t
	}
	// init paths
	paths := make([]int, n)
	for i := 0; i < n; i++ {
		paths[i] = math.MaxInt8
	}
	// start node set path to zero
	paths[k-1] = 0
	// init visited node
	visited := make([]bool, n)
	//  set k is min value, start from k
	// paths[k-1] = 0
	curNodeIdx := k - 1
	for curNodeIdx != -1 {
		for targetNodeIdx, curToTargetTime := range sourceTarget[curNodeIdx] {
			paths[targetNodeIdx] = min(paths[targetNodeIdx], paths[curNodeIdx]+curToTargetTime)
		}
		// record curNode is visited
		visited[curNodeIdx] = true

		// find min nodeIdx without visiting
		curNodeIdx = -1
		val := math.MaxInt8
		for i := 0; i < n; i++ {
			if !visited[i] && paths[i] < val {
				val = paths[i]
				curNodeIdx = i
			}
		}
	}
	val := -1
	for i := 0; i < n; i++ {
		if paths[i] > val {
			val = paths[i]
		}
	}
	if val == math.MaxInt8 {
		return -1
	}
	return val
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
