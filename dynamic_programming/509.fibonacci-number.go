package main

/*
 * @lc app=leetcode id=509 lang=golang
 *
 * [509] Fibonacci Number
 */

// @lc code=start
func fib(n int) int {
	if n < 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// @lc code=end
