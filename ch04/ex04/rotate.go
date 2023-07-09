package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4}
	s = rotate(s)
	fmt.Println(s)
	s = rotate(s)
	fmt.Println(s)
}

// [ "a", "b", "c", "d" ] => ["b", "c", "d", "a"]
func rotate(s []int) []int {
	return append(s[1:], s[0])
}
