package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	sort.Strings(strs)

	for i, char := range strs[0] {
		for _, str := range strs[1:] {
			char2 := rune(str[i])
			if char != char2 {
				return strs[0][:i]
			}

		}
	}
	return strs[0]
}

// @lc code=end
