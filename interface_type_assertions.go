package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string) // asserts that the interface value i holds the concrete type string and assigns the underlying string value to the variable s
	fmt.Println(s)

	s, ok := i.(string) // to test whether an interface value holds a specific type, a type assertion can return two values
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // as i does not hold a float64, this statement will trigger a panic
	fmt.Println(f)
}
