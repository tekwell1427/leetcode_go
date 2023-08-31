package main

/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	b := []byte(s)
	cur := len(b)
	for i := len(b) - 1; i >= 0; i-- {
		if b[i] == ' ' {
			cur--
		} else {
			break
		}
	}
	b = b[:cur]
	output := 0
	for i := len(b) - 1; i >= 0; i-- {
		if b[i] == ' ' {
			break
		}
		output++
	}
	return output
}

// @lc code=end
