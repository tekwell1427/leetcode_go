package main

/*
 * @lc app=leetcode id=739 lang=golang
 *
 * [739] Daily Temperatures
 */

// @lc code=start
func dailyTemperatures(temperatures []int) []int {
	output := make([]int, len(temperatures))
	stack := []int{}

	for i, temp := range temperatures {
		if len(stack) == 0 {
			stack = append(stack, i)
		} else {
			for len(stack) > 0 {
				topIdx := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if temperatures[topIdx] < temp {
					output[topIdx] = i - topIdx
				} else {
					stack = append(stack, topIdx)
					break
				}
			}
			stack = append(stack, i)
		}
	}
	return output
}

// @lc code=end
