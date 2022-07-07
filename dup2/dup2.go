package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	var filesWithRepeat []string
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			if isHasRepeatString(f) {
				filesWithRepeat = append(filesWithRepeat, f.Name())
			}
			countLines(f, counts)
			f.Close()
		}

	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", line, n)
		}
	}
	for _, v := range filesWithRepeat {
		fmt.Printf("%v\n", v)
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func isHasRepeatString(f *os.File) bool {
	input := bufio.NewScanner(f)
	strings := make(map[string]int)
	for input.Scan() {
		strings[input.Text()]++
		if strings[input.Text()] > 0 {
			return true
		}
	}

	return false
}
