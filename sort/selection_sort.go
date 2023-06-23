package main

import "fmt"

func SelectionSort(s []int) {
	for n := 0; n < len(s); n++ {
		minIdx := n
		for i := n + 1; i < len(s); i++ {
			if s[i] < s[minIdx] {
				minIdx = i
			}
		}
		tmp := s[n]
		s[n] = s[minIdx]
		s[minIdx] = tmp
	}
}

func main() {
	a := []int{1, 5, 7, 8, 4, 9, 10}
	SelectionSort(a)
	fmt.Println(a)
}
