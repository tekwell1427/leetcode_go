package main

/*
 * @lc app=leetcode id=187 lang=golang
 *
 * [187] Repeated DNA Sequences
 */

// @lc code=start
func findRepeatedDnaSequences(s string) []string {
	record := make(map[string]int)
	var output []string

	for i := 0; i < len(s)-9; i++ {
		cur := s[i : i+10]
		record[cur]++
		if record[cur] == 2 {
			output = append(output, cur)
		}
	}
	return output
}

// @lc code=end
