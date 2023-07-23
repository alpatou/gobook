package main

// package declaration , where it belongs
// main is special , defines a stand alone program, not a library

import (
	"bufio"
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

func dup1() {
	counts := make(map[string]int) // makes an empty map
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}
