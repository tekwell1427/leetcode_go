package main

/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Find the Index of the First Occurrence in a String
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	nLen := len(needle)
	prefix := getPrefix(needle)
	hIdx, prefixIdx := 0, 0
	output := -1
	for hIdx < len(haystack) {
		if haystack[hIdx] == needle[prefixIdx] {
			if prefixIdx == nLen-1 {
				output = hIdx - prefixIdx
				break
			}
			prefixIdx++
			hIdx++
		} else {
			if prefixIdx > 0 {
				prefixIdx = prefix[prefixIdx-1]
			} else {
				hIdx++
			}
		}
	}
	return output
}

func getPrefix(a string) []int {
	prefix := make([]int, len(a))
	l, r := 0, 1
	for r < len(a) {
		if a[l] == a[r] {
			l += 1
			prefix[r] = l
			r += 1
		} else {
			if l > 0 {
				l = prefix[l-1]
			} else {
				prefix[r] = 0
				r += 1
			}
		}
	}
	return prefix
}

// @lc code=end
