package main

/*
 * @lc app=leetcode id=1768 lang=golang
 *
 * [1768] Merge Strings Alternately
 */

// @lc code=start
func mergeAlternately(word1 string, word2 string) string {
	output := []byte{}

	i, j := 0, 0
	for ; i < len(word1) && j < len(word2); i, j = i+1, j+1 {
		output = append(output, word1[i], word2[j])
	}
	if i != len(word1) {
		output = append(output, word1[i:]...)
	}
	if j != len(word2) {
		output = append(output, word2[j:]...)
	}
	return string(output)
}

// @lc code=end
