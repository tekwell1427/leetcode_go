package main

import "fmt"

func BubbleSort(s []int) {
	for n := 0; n < len(s); n++ {
		for i := 0; i < len(s)-1; i++ {
			if s[i] > s[i+1] {
				tmp := s[i]
				s[i] = s[i+1]
				s[i+1] = tmp
			}
		}
	}
}

func main() {
	a := []int{1, 5, 7, 8, 4, 9, 10}
	BubbleSort(a)
	fmt.Println(a)
}
