package main

// package declaration , where it belongs
// main is special , defines a stand alone program, not a library

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	s, sep := "", ""
	for index, arg := range os.Args[0:] {
		s += strconv.Itoa(index) + sep + arg + "\n"
		sep = " "
	}
	fmt.Println(s)
	smartArgumentPrinter()
}

func smartArgumentPrinter() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func smartLoop() {
	for i, arg := range os.Args {
		fmt.Printf("%d: %s\n", i, arg)
	}
}
