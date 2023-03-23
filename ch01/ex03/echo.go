package main

import (
	"os"
	"strings"
)

func echo(args []string) string {
	return strings.Join(args, " ")
}

func slow_echo(args []string) string {
	s, sep := "", ""
	for _, str := range args {
		s += sep + str
		sep = " "
	}
	return s
}

func main() {
	echo(os.Args[1:])
	slow_echo(os.Args[1:])
}
