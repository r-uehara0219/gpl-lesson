package main

import "fmt"

func main() {
	s := []string{"a", "a", "b", "c", "c", "c", "d", "d", "e", "e", "e", "e"}
	s = removeAdjacentDuplicate(s)
	fmt.Println(s)
}

// []stringスライス内で隣接している重複を除去する
func removeAdjacentDuplicate(s []string) []string {
	var result []string

	for i, a := range s {
		if i == 0 {
			result = append(result, a)
			continue
		}

		if a != s[i-1] {
			result = append(result, a)
		}
	}

	return result
}
