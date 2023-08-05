package main

/*
 * @lc app=leetcode id=474 lang=golang
 *
 * [474] Ones and Zeroes
 */

// @lc code=start
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for dpRow := 0; dpRow < m+1; dpRow++ {
		dp[dpRow] = make([]int, n+1)
	}

	for _, str := range strs {
		zeroCount, oneCount := 0, 0
		for _, s := range str {
			if s == '0' {
				zeroCount++
			} else {
				oneCount++
			}
		}

		for curM := m; curM >= zeroCount; curM-- {
			for curN := n; curN >= oneCount; curN-- {
				dp[curM][curN] = max(dp[curM][curN], dp[curM-zeroCount][curN-oneCount]+1)
			}
		}
	}
	return dp[m][n]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
