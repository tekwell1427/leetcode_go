package main

/*
 * @lc app=leetcode id=345 lang=golang
 *
 * [345] Reverse Vowels of a String
 */

// @lc code=start
func reverseVowels(s string) string {
	b := []byte(s)
	left, right := 0, len(s)-1
	record := make(map[byte]bool)
	vowels := []byte{'a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U'}

	for _, vowel := range vowels {
		record[vowel] = true
	}
	for left < right {
		if !record[b[left]] {
			left++
			continue
		}
		if !record[b[right]] {
			right--
			continue
		}
		tmp := b[left]
		b[left] = b[right]
		b[right] = tmp
		left++
		right--
	}
	return string(b)
}

// @lc code=end
