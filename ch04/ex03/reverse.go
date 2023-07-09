package main

import "fmt"

const Size = 10

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s = reverse(s)
	fmt.Println(s)

	a := [Size]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	reverseToPointer(&a)
	fmt.Println(a)
}

func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func reverseToPointer(s *[Size]int) {
	for i, j := 0, Size-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
