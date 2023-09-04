package main

import "fmt"

func InsertionSort(s []int) int {
	count := 0
	// for n := len(s) - 1; n >= 0; n-- {
	for n := 0; n < len(s); n++ {
		cur := s[n]
		for i := n - 1; i >= 0; i-- {
			count++
			if cur < s[i] {
				s[i+1] = s[i]
				s[i] = cur
			} else {
				break
			}
		}
	}
	return count
}

func main() {
	a := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	count := InsertionSort(a)
	fmt.Println("worst case: ", a, "count:", count)
	a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	count = InsertionSort(a)
	fmt.Println("best case: ", a, "count:", count)
}
