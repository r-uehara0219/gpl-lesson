package main

import "bytes"

func main() {
	println(comma("1234567890"))
	println(commaWithBuffer("1234567890"))
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func commaWithBuffer(s string) string {
	var buf bytes.Buffer
	bytes := []byte(s)
	for i, b := range bytes {
		if i != 0 && (len(bytes)-i)%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(b)
	}

	return buf.String()
}
