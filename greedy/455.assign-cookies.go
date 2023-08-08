package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=455 lang=golang
 *
 * [455] Assign Cookies
 */

// @lc code=start
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	numChild := 0
	for _, cookie := range s {

		if cookie >= g[numChild] {
			numChild++
			if numChild == len(g) {
				break
			}
		}
	}

	return numChild
}

// @lc code=end
