package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	s := []byte("a  b  c  d  あ  f  g  h  i  j  k  l  m  n  o  p")
	s = squashSpaces(s)
	fmt.Println(string(s))
}

// UTF-8でエンコードされた[]byteスライス内で隣接しているUnicodeスペースを、ASCIIスペースに置き換える
func squashSpaces(s []byte) []byte {
	var result []byte
	const space = byte(' ')
	beforeSpace := false // 直前の文字がスペースかどうか
	offset := 0          // 現在処理しているUTF-8文字の先頭バイトからのオフセット

	for i := range s {
		r, _ := utf8.DecodeRune(s[i-offset : i+1])
		if r == utf8.RuneError {
			offset++
			continue
		}

		if unicode.IsSpace(r) {
			if !beforeSpace {
				result = append(result, space)
			}
			beforeSpace = true
		} else {
			result = append(result, s[i-offset:i+1]...)
			beforeSpace = false
		}

		offset = 0
	}

	return result
}
