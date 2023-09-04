package main

import "fmt"

func BubbleSort(s []int) int {
	count := 0
	for n := 0; n < len(s); n++ {
		ischange := false
		for i := 0; i < len(s)-1; i++ {
			count++
			if s[i] > s[i+1] {
				tmp := s[i]
				s[i] = s[i+1]
				s[i+1] = tmp
				ischange = true
			}
		}
		if !ischange {
			break
		}
	}
	return count
}

func main() {
	a := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	count := BubbleSort(a)
	fmt.Println("worst case: ", a, "count:", count)
	a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	count = BubbleSort(a)
	fmt.Println("best case: ", a, "count:", count)
}
