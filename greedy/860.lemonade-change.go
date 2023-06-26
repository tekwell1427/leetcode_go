package main

/*
 * @lc app=leetcode id=860 lang=golang
 *
 * [860] Lemonade Change
 */

// @lc code=start
func lemonadeChange(bills []int) bool {
	hands := make(map[int]int)
	for _, bill := range bills {
		hands[bill]++
		change := bill - 5
		if change == 5 {
			if hands[5] == 0 {
				return false
			}
			hands[5]--
		}
		if change == 15 {
			if hands[10] > 0 && hands[5] > 0 {
				hands[10]--
				hands[5]--
			} else if hands[10] == 0 && hands[5] >= 3 {
				hands[5] -= 3
			} else {
				return false
			}
		}
	}
	return true
}

// @lc code=end
