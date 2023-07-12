package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode id=394 lang=golang
 *
 * [394] Decode String
 */

// @lc code=start
func decodeString(s string) string {
	output := ""
	numStack := []int{}
	stringStack := []string{}
	multi, res := 0, ""

	for _, char := range s {
		if char >= '0' && char <= '9' {
			stringStack = append(stringStack, res)
			res = ""

			c, _ := strconv.Atoi(string(char))
			multi = multi*10 + c
		} else if char == '[' {
			numStack = append(numStack, multi)
			multi = 0
		} else if char == ']' {
			num := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]

			output += strings.Repeat(output, num)
			res = ""

			if len(stringStack) > 0 {
				s := stringStack[len(stringStack)-1]
				stringStack = stringStack[:len(stringStack)-1]
				output = s + output
			}
			fmt.Println(output, "YY")
		} else {
			res += string(char)
		}
		// fmt.Println(char, stringStack, numStack)
	}

	return output
}

// @lc code=end
