package main

/*
 * @lc app=leetcode id=96 lang=golang
 *
 * [96] Unique Binary Search Trees
 */

// @lc code=start
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		sum := 0
		for j := i; j > 0; j-- {
			// left+right = i-1, and right= i-j
			left, right := j-1, i-j
			sum += dp[left] * dp[right]
		}
		dp[i] = sum
	}
	return dp[n]
}

// @lc code=end
