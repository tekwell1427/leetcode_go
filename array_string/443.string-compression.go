package main

import (
	"strconv"
)

/*
 * @lc app=leetcode id=443 lang=golang
 *
 * [443] String Compression
 */

// @lc code=start
func compress(chars []byte) int {
	count := 0
	curIdx := 0
	for i, _ := range chars {
		count++
		if i == len(chars)-1 || chars[i] != chars[i+1] {
			chars[curIdx] = chars[i]
			curIdx++
			if count > 1 {
				countToStr := strconv.Itoa(count)
				l := len(countToStr)
				for i, j := curIdx, 0; i < curIdx+l; i, j = i+1, j+1 {
					chars[i] = countToStr[j]
				}
				curIdx += l
			}
			count = 0
		}
	}

	return curIdx
}

// @lc code=end
