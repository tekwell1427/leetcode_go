package main

/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 */

// @lc code=start
func isPalindrome(s string) bool {
	b := []byte(s)
	cur := 0
	for _, c := range b {
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
			if c >= 'A' && c <= 'Z' {
				c = c + ('a' - 'A')
			}
			b[cur] = c
			cur++
		}
	}
	b = b[:cur]
	left, right := 0, len(b)-1
	for left < right {
		if b[left] != b[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// @lc code=end
