package main

/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	// strs[i] consists of lowercase English letters.
	record := make(map[[26]int][]string)

	for _, str := range strs {
		sMap := [26]int{}
		for _, s := range str {
			sMap[s-'a']++
		}
		record[sMap] = append(record[sMap], str)
	}
	output := [][]string{}
	for _, value := range record {
		output = append(output, value)
	}
	return output
}

// @lc code=end
