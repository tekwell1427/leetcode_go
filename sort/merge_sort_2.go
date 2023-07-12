package main

import "fmt"

func MergeSort(s []int) {
		if len(s)<2{
			return 
		}
		left, right:=0, len(s)-1
		mid := (left + right) / 2
		helper(s[:mid])
		helper(s[mid:])
		
		leftStart, rightStart:= 0, mid

		for leftStart<mid&&rightStart<len(s) {
			if s[leftStart] > s[rightStart] {
				output = append(output, r[j])
				j++
			} else {
				output = append(output, l[i])
				i++
			}
		}
		for ; i < len(l); i++ {
			output = append(output, l[i])
		}
		for ; j < len(r); j++ {
			output = append(output, r[j])
		}
		return output
	}

}

func main() {
	a := []int{1, 5, 7, 8, 4, 9, 10}
	MergeSort(a)
	fmt.Println(a)
}
