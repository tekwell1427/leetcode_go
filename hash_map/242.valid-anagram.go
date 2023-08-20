package main

/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	record1 := [26]int{}
	record2 := [26]int{}

	for _, c := range s {
		record1[c-'a']++
	}

	for _, c := range t {
		record2[c-'a']++
	}

	for i := 0; i < len(record1); i++ {
		if record1[i] != record2[i] {
			return false
		}
	}
	return true
}

// @lc code=end
