package main

/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */

// @lc code=start
func climbStairs(n int) int {
	dp := make([]int, n+1)
	steps := []int{1, 2}
	dp[0] = 1

	for curN := 0; curN <= n; curN++ {
		for _, curStep := range steps {
			if curN >= curStep {
				dp[curN] += dp[curN-curStep]
			}
		}
	}
	return dp[n]
}

// @lc code=end
