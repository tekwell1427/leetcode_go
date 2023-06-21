package main

import "sort"

/*
 * @lc app=leetcode id=455 lang=golang
 *
 * [455] Assign Cookies
 */

// @lc code=start
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	curChidIdx := 0
	for _, cookieSize := range s {

		if cookieSize >= g[curChidIdx] {
			curChidIdx++
			if curChidIdx == len(g) {
				break
			}
		}
	}
	return curChidIdx
}

// @lc code=end
