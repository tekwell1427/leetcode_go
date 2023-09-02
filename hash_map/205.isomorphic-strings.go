package main

/*
 * @lc app=leetcode id=205 lang=golang
 *
 * [205] Isomorphic Strings
 */

// @lc code=start
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	// because is ascii
	recordS := [128]byte{}
	recordT := [128]byte{}

	for i := 0; i < len(s); i++ {
		if recordS[s[i]] == 0 && recordT[t[i]] != 0 {
			return false
		}
		if recordS[s[i]] != 0 && recordT[t[i]] == 0 {
			return false
		}
		if recordS[s[i]] == 0 && recordT[t[i]] == 0 {
			recordS[s[i]] = t[i]
			recordT[t[i]] = s[i]
		}
		if recordS[s[i]] != t[i] || recordT[t[i]] != s[i] {
			return false
		}
	}
	return true
}

// @lc code=end
