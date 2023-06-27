package main

import (
	"strconv"
)

/*
 * @lc app=leetcode id=738 lang=golang
 *
 * [738] Monotone Increasing Digits
 */

// @lc code=start
func monotoneIncreasingDigits(n int) int {
	s := strconv.Itoa(n)
	output := make([]int, len(s))
	preNum := s[0]
	output[0] = int(preNum) - 48
	for i := 1; i < len(s); i++ {
		curNum := s[i]
		if curNum >= preNum {
			curInt := int(curNum) - 48
			output[i] = curInt
		} else {
			j := i - 1
			for ; j > 0 && output[j]-1 < output[j-1]; j-- {

			}
			output[j]--
			for z := j + 1; z < len(s); z++ {
				output[z] = 9
			}
			break
		}
		preNum = curNum
	}
	out := ""
	for i := 0; i < len(output); i++ {
		out += strconv.Itoa(output[i])
	}
	o, _ := strconv.Atoi(out)
	return o
}

// @lc code=end
