package main

import (
	"strings"
)

/*
 * @lc app=leetcode id=151 lang=golang
 *
 * [151] Reverse Words in a String
 */

// @lc code=start
func reverseWords(s string) string {
	b := strings.Split(strings.Trim(s, " "), " ")
	left, right := 0, len(b)-1

	for left < right {
		if b[left] == " " {
			left++
			continue
		}
		if b[right] == " " {
			right--
			continue
		}
		tmp := b[left]
		b[left] = b[right]
		b[right] = tmp
		left++
		right--
	}
	return strings.Join(b, " ")
}

// @lc code=end
