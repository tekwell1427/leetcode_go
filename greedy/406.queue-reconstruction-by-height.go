package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=406 lang=golang
 *
 * [406] Queue Reconstruction by Height
 */

// @lc code=start
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	for i, p := range people {
		insertIdx := p[1]
		copy(people[insertIdx+1:i+1], people[insertIdx:i])
		people[insertIdx] = p
	}
	return people
}

// @lc code=end
