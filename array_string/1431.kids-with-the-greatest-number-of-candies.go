package main

/*
 * @lc app=leetcode id=1431 lang=golang
 *
 * [1431] Kids With the Greatest Number of Candies
 */

// @lc code=start
func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandy := 0
	for _, candy := range candies {
		maxCandy = max(maxCandy, candy)
	}
	output := []bool{}
	for _, candy := range candies {
		var isMax bool = false
		if candy+extraCandies >= maxCandy {
			isMax = true
		}
		output = append(output, isMax)
	}
	return output
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
