package main

import "fmt"

func InsertionSort(s []int) {
	for n := 1; n < len(s); n++ {
		for i := n - 1; i >= 0; i-- {
			if s[i+1] < s[i] {
				tmp := s[i+1]
				s[i+1] = s[i]
				s[i] = tmp
			}
		}
	}
}

func main() {
	a := []int{1, 5, 7, 8, 4, 9, 10}
	InsertionSort(a)
	fmt.Println(a)
}
