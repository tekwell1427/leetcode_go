package main

import (
	"math"
	"strconv"
)

/*
 * @lc app=leetcode id=202 lang=golang
 *
 * [202] Happy Number
 */

// @lc code=start
func isHappy(n int) bool {
	record := make(map[int]bool)

	for {
		s := strconv.Itoa(n)
		newN := 0
		for _, c := range s {
			i, _ := strconv.Atoi(string(c))
			newN += int(math.Pow(float64(i), float64(2)))
		}
		if newN == 1 {
			return true
		}
		if _, exits := record[newN]; exits {
			return false
		}
		record[newN] = true
		n = newN
	}
}

// @lc code=end
