package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma("1234567890"))
	fmt.Println(comma("-1234567890"))
	fmt.Println(comma("12345.67890"))
}

// inserts commas
func comma(s string) string {
	if s[0] == '+' || s[0] == '-' {
		return string(s[0]) + comma(s[1:])
	}

	if i := strings.Index(s, "."); i >= 0 {
		return comma(s[:i]) + s[i:]
	}

	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
