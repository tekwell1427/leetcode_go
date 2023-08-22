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
		t, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, t)
		} else {
			num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
			switch token {
			case "+":
				stack[len(stack)-2] = num1 + num2
			case "-":
				stack[len(stack)-2] = num1 - num2
			case "*":
				stack[len(stack)-2] = num1 * num2
			case "/":
				stack[len(stack)-2] = num1 / num2
			}
			stack = stack[:len(stack)-1]
		}
	}
	return stack[0]
}

// @lc code=end
