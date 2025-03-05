package main

import "fmt"

func main() {
	// Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
	defer fmt.Println("world")
	defer fmt.Println("cold")

	fmt.Println("hello")
}

