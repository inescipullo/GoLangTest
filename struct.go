package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

// Defining Struct Literals
var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
	// Accessing Struct Fields
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v)
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

