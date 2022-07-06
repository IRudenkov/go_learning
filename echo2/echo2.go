package main

import (
	"fmt"
	"os"
	"strconv"
)

//second version echo, with loop for range

func main() {
	for index, args := range os.Args[1:] {
		fmt.Println(strconv.Itoa(index+1) + " " + args)
	}

}
