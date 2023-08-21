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
	// cheat
	s = strings.Trim(s, " ")
	b := []byte(s)
	// remove continuous empty
	left := 0
	for i, c := range b {
		if c == ' ' && b[i-1] == ' ' {
			continue
		}
		b[left] = c
		left++
	}
	// remove redundancy element
	b = b[:left]

	// reverse (character-wise)
	left, right := 0, len(b)-1
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
	// reverse word-wise
	for i := 0; i < len(b); i++ {
		j := i
		for ; j < len(b) && b[j] != ' '; j++ {
		}
		subS := b[i:j]
		left, right := 0, j-i-1
		for left < right {
			subS[left], subS[right] = subS[right], subS[left]
			left++
			right--
		}
		i = j
	}
	return string(b)
}

// @lc code=end
