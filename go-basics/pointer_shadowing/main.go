package main

import (
	"fmt"
	"unsafe"
)

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

	var ui1 uint16
	fmt.Printf("memory addreos of i: %p\n", &ui1)
	var ui2 uint16
	fmt.Printf("memory addreos of i: %p\n", &ui2)

	var p1 *uint16
	fmt.Printf("value of p1: %v\n", p1)
	p1 = &ui1
	fmt.Printf("value of p1: %v\n", p1)
	fmt.Printf("size of p1: %d[bytes]\n", unsafe.Sizeof(p1))
	fmt.Printf("memory address of p1: %p\n", &p1)
	fmt.Printf("value of ui1(dereference): %v\n", *p1)
	*p1 = 1
	fmt.Printf("value of ui1: %v\n\n", ui1)

	var pp1 **uint16 = &p1
	fmt.Printf("value of pp1: %v\n", pp1)
	fmt.Printf("memory address of pp1: %p\n", &pp1)
	fmt.Printf("size of pp1: %d[bytes]\n", unsafe.Sizeof(pp1))
	fmt.Printf("value of p1(dereference): %v\n", *pp1)
	fmt.Printf("value of ui1(double dereference): %v\n", **pp1)
	**pp1 = 10
	fmt.Printf("value of ui1: %v\n\n", ui1)

	ok, result := true, "A"
	fmt.Printf("memory address of result: %p\n", &result)
	if ok {
		// result := "B" // shadowing(外側のresultと内側のresultは別物で、ポインタも別物)
		result = "B"
		fmt.Printf("memory address of result: %p\n", &result)
		println(result)
	} else {
		// result := "C"
		result = "C"
		fmt.Printf("memory address of result: %p\n", &result)
		println(result)
	}
	println(result)
}
