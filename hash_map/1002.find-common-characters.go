package main

/*
 * @lc app=leetcode id=1002 lang=golang
 *
 * [1002] Find Common Characters
 */

// @lc code=start
func commonChars(words []string) []string {
	record := [26]int{}

	for _, c := range words[0] {
		record[c-'a']++
	}

	for i := 1; i < len(words); i++ {
		word := words[i]
		record2 := [26]int{}
		for _, c := range word {
			record2[c-'a']++
		}
		record = min(record, record2)
	}

	output := []string{}
	for i, c := range record {
		for c > 0 {
			output = append(output, string(i+int('a')))
			c--
		}
	}
	return output
}
func min(a [26]int, b [26]int) [26]int {
	for i := 0; i < 26; i++ {
		var result int
		if a[i] > b[i] {
			result = b[i]
		} else {
			result = a[i]
		}
		a[i] = result
	}
	return a
}

// @lc code=end
