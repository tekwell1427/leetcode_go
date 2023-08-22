package main

/*
 * @lc app=leetcode id=1047 lang=golang
 *
 * [1047] Remove All Adjacent Duplicates In String
 */

// @lc code=start
func removeDuplicates(s string) string {
	stack := []rune{}

	for _, c := range s {
		if len(stack) != 0 {
			if stack[len(stack)-1] != c {
				stack = append(stack, c)
			} else {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, c)
		}
	}
	return string(stack)
}

// @lc code=end
