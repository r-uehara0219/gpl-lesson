// $ ./ex02 10 13.1
// 10尺 = 33メートル, 10メートル = 3.03尺
// 13.1尺 = 43.2メートル, 13.1メートル = 3.97尺

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, arg := range getArgs() {
		length, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "shaku_meter: %v\n", err)
			os.Exit(1)
		}
		s := Shaku(length)
		m := Meter(length)
		fmt.Printf("%s = %s, %s = %s\n",
			s, ShakuToMeter(s), m, MeterToShaku(m))
	}
}

func getArgs() []string {
	args := os.Args[1:]
	if len(args) == 0 {
		// Ctrl+Dで入力を終了する
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			args = append(args, input.Text())
		}
	}
	return args
}
