package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=1657 lang=golang
 *
 * [1657] Determine if Two Strings Are Close
 */

// @lc code=start
func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	// only lowercase, slice for sort
	record1 := make([]int, 26)
	record2 := make([]int, 26)

	for i := 0; i < len(word1); i++ {
		record1[word1[i]-'a']++
		record2[word2[i]-'a']++
	}

	for i := 0; i < 26; i++ {
		if (record1[i] > 0 && record2[i] == 0) || (record1[i] == 0 && record2[i] > 0) {
			return false
		}
	}
	sort.Ints(record1)
	sort.Ints(record2)

	for i := 0; i < len(record1); i++ {
		if record1[i] != record2[i] {
			return false
		}
	}
	return true
}

// @lc code=end
