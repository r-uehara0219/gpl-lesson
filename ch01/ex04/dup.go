// 出現した単語・出現回数・出現したファイル名を出力する

// $ go run dup.go text/example1.txt text/example2.txt
// 3	apple
// text/example1.txt text/example2.txt
// 3	banana
// text/example1.txt text/example2.txt
// 2	cat
// text/example1.txt text/example2.txt
// 6	dictionary
// text/example1.txt text/example2.txt
// 2	end
// text/example1.txt

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
		counts := make(map[string]int)
	detected_files := make(map[string]map[string]bool)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, detected_files)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, detected_files)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			for file_name, _ := range detected_files[line] {
				fmt.Printf("%s ", file_name)
			}
			fmt.Printf("\n")
		}
	}
}

func countLines(f *os.File, counts map[string]int, detected_files map[string]map[string]bool) {
	input := bufio.NewScanner(f)
	file_name := f.Name()
	for input.Scan() {
		counts[input.Text()]++
		if detected_files[input.Text()] == nil {
			detected_files[input.Text()] = make(map[string]bool)
		}
		detected_files[input.Text()][file_name] = true
	}
	// NOTE: ignoring potential errors from input.Err()
}
