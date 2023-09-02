package main

import "strings"

/*
 * @lc app=leetcode id=290 lang=golang
 *
 * [290] Word Pattern
 */

// @lc code=start
func wordPattern(pattern string, s string) bool {
	newS := strings.Split(s, " ")
	if len(newS) != len(pattern) {
		return false
	}
	recordP := make(map[byte]string)
	recordS := make(map[string]byte)

	for i := 0; i < len(pattern); i++ {
		p := pattern[i]
		s := newS[i]
		if value, exists := recordP[p]; exists {
			if value != newS[i] {
				return false
			}
		}
		if value, exists := recordS[s]; exists {
			if value != p {
				return false
			}
		}
		recordP[p] = s
		recordS[s] = p
	}
	return true
}

// @lc code=end
