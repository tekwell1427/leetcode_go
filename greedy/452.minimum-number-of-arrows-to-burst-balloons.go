package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=452 lang=golang
 *
 * [452] Minimum Number of Arrows to Burst Balloons
 */

// @lc code=start
func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	output := 1
	shotEnd := points[0][1]
	for i := 1; i < len(points); i++ {
		curStart, curEnd := points[i][0], points[i][1]
		if curStart <= shotEnd {
			shotEnd = min(shotEnd, curEnd)
		} else if curStart > shotEnd {
			output++
			shotEnd = curEnd
		}
	}
	return output
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end
