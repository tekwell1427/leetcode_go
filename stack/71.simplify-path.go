package main

import (
	"strings"
)

/*
 * @lc app=leetcode id=71 lang=golang
 *
 * [71] Simplify Path
 */

// @lc code=start
func simplifyPath(path string) string {
	stack := []string{}
	// p := []byte(path)
	p := strings.Split(path, "/")

	for _, e := range p {
		if e == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if e == "" || e == "." {
			continue
		} else {
			stack = append(stack, e)
		}
	}
	return "/" + strings.Join(stack, "/")
}

// @lc code=end
