package main

/*
 * @lc app=leetcode id=2390 lang=golang
 *
 * [2390] Removing Stars From a String
 */

// @lc code=start
func removeStars(s string) string {
	charStack := []byte{}

	for i := 0; i < len(s); i++ {
		char := s[i]
		if char == '*' {
			charStack = charStack[:len(charStack)-1]
		} else {
			charStack = append(charStack, byte(char))
		}
	}
	return string(charStack)
}

// @lc code=end
