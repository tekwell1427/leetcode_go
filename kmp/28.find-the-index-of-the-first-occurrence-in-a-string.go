package main

/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Find the Index of the First Occurrence in a String
 */

// @lc code=start
var prefixTable []int

func strStr(haystack string, needle string) int {

	return kmpSearch(haystack, needle)
}
func getPrefix(a string) []int {
	prefix := make([]int, len(a))

	l, i := 0, 1
	for i < len(a) {
		if a[i] == a[l] {
			l++
			prefix[i] = l
			i++
		} else {
			if l > 0 {
				l = prefix[l-1]
				// l 此時是0
			} else {
				prefix[i] = l
				i++
				// 不能i++，因為a[i]!=a[l]
			}
		}
	}
	return prefix
}

// 往右移動
func movePrefix(prefix []int) {
	for i := len(prefix) - 1; i > 0; i-- {
		prefix[i] = prefix[i-1]
	}
	prefix[0] = -1
}

func kmpSearch(input string, pattern string) int {
	output := -1
	prefixTable := getPrefix(pattern)
	movePrefix(prefixTable)

	i, j := 0, 0
	for i < len(input) {
		if input[i] == pattern[j] {
			if j == len(pattern)-1 {
				output = i - j
				break
			}
			i++
			j++
		} else {
			j = prefixTable[j]
			// 如果pattern開頭跟input開頭 是不一樣的，則都要往下走
			if j == -1 {
				i++
				j++
			}
		}
	}
	return output
}

// @lc code=end
