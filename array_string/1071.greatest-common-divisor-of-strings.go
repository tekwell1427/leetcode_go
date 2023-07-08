package main

/*
 * @lc app=leetcode id=1071 lang=golang
 *
 * [1071] Greatest Common Divisor of Strings
 */

// @lc code=start
func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	g := GCD(len(str1), len(str2))
	return string(str1[:g])
}
func GCD(a, b int) int {
	for {
		r := a % b
		if r == 0 {
			return b
		}
		a = b
		b = r
	}
}

// @lc code=end
