package main

/*
 * @lc app=leetcode id=416 lang=golang
 *
 * [416] Partition Equal Subset Sum
 */

// @lc code=start
func canPartition(nums []int) bool {
	var sum int
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}
	mean := sum / 2

	dp := make([]int, mean+1)
	for objectIdx := 0; objectIdx < len(nums); objectIdx++ {
		num := nums[objectIdx]
		for j := mean; j >= num; j-- {
			if dp[j-num]+num <= mean {
				dp[j] = max(dp[j], dp[j-num]+num)
			}
		}
	}
	return dp[len(dp)-1] == mean
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
