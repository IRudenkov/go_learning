package main

import (
	"fmt"
	"os"
)

//second version echo, with loop for range

func main() {
	s, sep := "", ""
	for _, args := range os.Args[1:] {
		s += sep + args
		sep = " "
	}
	fmt.Println(s)
}
