package main

import "fmt"

const secret = "abc"

type Os int 

const (
	Mac Os = iota + 1
	Windows
	Linux
)

var (
	s string
	i int
	f float64
	b bool
	os Os
)

func main() {
	// var i int 
	// var i int = 2
	// var i = 2
	i := 1
	ui := uint16(2)
	fmt.Println(ui)
	fmt.Printf("i: %v %T\n", i, i)
	fmt.Printf("i: %[1]v %[1]T ui: %[2]v %[2]T\n", i, ui)

	f := 1.23456
	s := "hello"
	b := true

	fmt.Printf("f: %[1]v %[1]T\n", f)
	fmt.Printf("s: %[1]v %[1]T\n", s)
	fmt.Printf("b: %[1]v %[1]T\n", b)

	pi, tille := 3.14, "Go"
	fmt.Printf("pi: %[1]v %[1]T tille: %[2]v %[2]T\n", pi, tille)

	x := 10
	y := 20.2
	z := float64(x) + y
	fmt.Printf("z: %[1]v %[1]T\n", z)

	fmt.Printf("Mac: %v, Windows: %v, Linux: %v\n", Mac, Windows, Linux)
}
