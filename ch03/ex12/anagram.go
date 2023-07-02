package main

import (
	"fmt"
)

func main() {
	fmt.Println(isAnagram("abc", "cba"))
	fmt.Println(isAnagram("abc", "cab"))
	fmt.Println(isAnagram("안녕하세요", "요세하녕안"))
	fmt.Println(isAnagram("안녕하세요", "세하녕안안"))
}

func isAnagram(s1 string, s2 string) bool {
	r1 := []rune(s1)
	r2 := []rune(s2)
	if len(r1) != len(r2) {
		return false
	}

	for i := 0; i < len(r1); i++ {
		if r1[i] != r2[len(r2)-1-i] {
			return false
		}
	}
	return true
}
