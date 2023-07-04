package main

/*
 * @lc app=leetcode id=1004 lang=golang
 *
 * [1004] Max Consecutive Ones III
 */
// 毛毛蟲跑法，頭動，尾巴固定；尾巴動，頭固定
// @lc code=start
func longestOnes(nums []int, k int) int {
	left, right := 0, 0
	output := 0
	zeroCount := 0
	for ; right < len(nums); right++ {
		num := nums[right]
		if num == 0 {
			zeroCount++
		}
		for zeroCount > k {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}
		output = max(output, right-left+1)
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
