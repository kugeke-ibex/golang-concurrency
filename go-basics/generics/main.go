package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)


type customConstraint interface {
	~int | int16 | float32 | float64 | string
}
type Newint int

func add[T customConstraint](a, b T) T {
	return a + b
}

func min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func sumValues[K int | string, V constraints.Float | constraints.Integer] (m map[K]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}
	return sum
 }


func main() {
	fmt.Printf("%v\n", add(1, 2))
	fmt.Printf("%v\n", add(1.1, 2.2))
	fmt.Printf("%v\n", add("Hello", "World"))
	var i1, i2 Newint = 3, 4
	fmt.Printf("%v\n", add(i1, i2))
	fmt.Printf("%v\n", min(i1, i2))
	m1 := map[string]uint{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	m2 := map[string]float32{
		"a": 1.23,
		"b": 4.56,
		"c": 7.89,
	}
	fmt.Printf("%v\n", sumValues(m1))
	fmt.Printf("%v\n", sumValues(m2))
}