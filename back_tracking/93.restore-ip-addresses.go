package main

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode id=93 lang=golang
 *
 * [93] Restore IP Addresses
 */

// @lc code=start
func restoreIpAddresses(s string) []string {
	output := []string{}
	var helper func(preIdx int, preString string, dotNum int)
	helper = func(preIdx int, preString string, dotNum int) {
		if dotNum == 4 {
			if preIdx == len(s)-1 {
				// 把dot拿掉
				tmp := strings.Clone(preString)
				tmp = tmp[:len(tmp)-1]
				output = append(output, tmp)
			}
			return
		}
		curS := ""
		for i := preIdx + 1; i < len(s); i++ {
			curS = curS + string(s[i])
			if i > preIdx+1 && s[preIdx+1] == '0' {
				return
			}
			curI, _ := strconv.Atoi(curS)
			if curI >= 0 && curI <= 255 {
				preString += curS + "."
				dotNum++
				helper(i, preString, dotNum)
				preString = preString[:len(preString)-len(curS)-1]
				dotNum--
			}
		}
	}
	helper(-1, "", 0)
	return output
}

// @lc code=end
