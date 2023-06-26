package main

import "sort"

/*
 * @lc app=leetcode id=435 lang=golang
 *
 * [435] Non-overlapping Intervals
 */

// @lc code=start
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	output := 0
	preEndVal := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		curStart, curEnd := intervals[i][0], intervals[i][1]
		// overlap
		if curStart < preEndVal {
			output++
			// remove interval which has larger endVal
			preEndVal = min(preEndVal, curEnd)
		} else {
			preEndVal = curEnd
		}
	}
	return output
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end
