package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	echo1()
	echo2()
	echo3()
	echo4()
}

func echo1() {
	var s, sep string  // If a variable is not explicitly initialized, it's implicitly initialized to the zero value for its type
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	s, sep := "", ""  // To declare explicitly initialized variables
	for _, arg := range os.Args[1:] {  // Loops through the values or os.Args, yielding a tuple of indexes and elements
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))  
	// When getting the subsequence of a slice, if either end of the range if omitted, it defaults to 0 or len respectively
}

func echo4() {
	fmt.Println(os.Args[1:])  // If we want to let Println format the output for us (uses Join, but adds brackets)
}
