package main

import "sort"

/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	output := [][]int{}
	preStart, preEnd := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		curStart, curEnd := intervals[i][0], intervals[i][1]
		if curStart > preEnd {
			output = append(output, []int{preStart, preEnd})
			preStart, preEnd = curStart, curEnd
		} else {
			preEnd = max(preEnd, curEnd)
		}
	}
	output = append(output, []int{preStart, preEnd})
	return output
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
