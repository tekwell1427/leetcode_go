package main

/*
 * @lc app=leetcode id=134 lang=golang
 *
 * [134] Gas Station
 */

// @lc code=start
func canCompleteCircuit(gas []int, cost []int) int {
	output := -1
	curSum := 0
	totalSum := 0
	for i, _ := range gas {
		excess := gas[i] - cost[i]
		curSum += excess
		totalSum += excess
		if curSum <= 0 && i != len(gas)-1 {
			curSum = 0
			output = -1
		} else if output == -1 {
			output = i
		}
	}
	if totalSum < 0 {
		return -1
	}
	return output
}

// @lc code=end
