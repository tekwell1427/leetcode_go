package main

import "sort"

/*
 * @lc app=leetcode id=274 lang=golang
 *
 * [274] H-Index
 */

// @lc code=start
func hIndex(citations []int) int {
	sort.Ints(citations)
	h := len(citations)

	for i := 0; i < len(citations); i++ {
		citation := citations[i]
		if citation >= h {
			break
		}
		h--
	}
	return h
}

// @lc code=end
