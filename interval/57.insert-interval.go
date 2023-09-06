package main

/*
 * @lc app=leetcode id=57 lang=golang
 *
 * [57] Insert Interval
 */

// @lc code=start
func insert(intervals [][]int, newInterval []int) [][]int {
	output := [][]int{}
	i := 0
	for i < len(intervals) && intervals[i][1] < newInterval[0] {
		output = append(output, intervals[i])
		i++
	}
	for i < len(intervals) && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(intervals[i][0], newInterval[0])
		newInterval[1] = max(intervals[i][1], newInterval[1])
		i++
	}
	output = append(output, newInterval)
	for i < len(intervals) {
		output = append(output, intervals[i])
		i++
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
	if a < b {
		return a
	}
	return b
}

// @lc code=end
