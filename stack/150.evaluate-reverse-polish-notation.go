package main

import "strconv"

/*
 * @lc app=leetcode id=150 lang=golang
 *
 * [150] Evaluate Reverse Polish Notation
 */

// @lc code=start
func evalRPN(tokens []string) int {
	stack := []int{}

	for _, token := range tokens {
		if token == "+" {
			result := stack[len(stack)-2] + stack[len(stack)-1]
			stack[len(stack)-2] = result
			stack = stack[:len(stack)-1]
		} else if token == "-" {
			result := stack[len(stack)-2] - stack[len(stack)-1]
			stack[len(stack)-2] = result
			stack = stack[:len(stack)-1]
		} else if token == "*" {
			result := stack[len(stack)-2] * stack[len(stack)-1]
			stack[len(stack)-2] = result
			stack = stack[:len(stack)-1]
		} else if token == "/" {
			result := stack[len(stack)-2] / stack[len(stack)-1]
			stack[len(stack)-2] = result
			stack = stack[:len(stack)-1]
		} else {
			t, _ := strconv.Atoi(token)
			stack = append(stack, t)
		}
	}
	return stack[0]
}

// @lc code=end
