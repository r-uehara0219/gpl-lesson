package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	c := input.Text()

	flag.Parse()
	args := flag.Args()

	switch args[0] {
	case "256":
		fmt.Printf("%x\n", sha256.Sum256([]byte(c)))
	case "384":
		fmt.Printf("%x\n", sha512.Sum384([]byte(c)))
	case "512":
		fmt.Printf("%x\n", sha512.Sum512([]byte(c)))
	default:
		fmt.Printf("%x\n", sha256.Sum256([]byte(c)))
	}
}
