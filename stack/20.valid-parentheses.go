package main

/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
func isValid(s string) bool {
	stack := []rune{}

	for _, n := range s {
		if n == ')' {
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		} else if n == ']' {
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		} else if n == '}' {
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, n)
		}
	}
	return len(stack) == 0
}

// @lc code=end
