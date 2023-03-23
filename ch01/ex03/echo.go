package main

import (
	"fmt"
	"os"
)

func echo() {
	fmt.Println(os.Args)
}

func slow_echo() {
	s, sep := "", ""
	for _, str := range os.Args[1:] {
		s += sep + str
		sep = " "
	}
	fmt.Println(s)
}

func main() {
	echo()
	slow_echo()
}
