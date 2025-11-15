package main

import (
	"fmt"
	"time"
)

func main() {
	a := -1
	if a == 0 {
		fmt.Println("a is 0")
	} else if a > 0 {
		fmt.Println("a is positive")
	} else {
		fmt.Println("a is negative")
	}
	for i := 0; i < 5; i++ {
		fmt.Printf("i: %v\n", i)
	}
	// for  {
	// 	fmt.Println("loop")
	// 	time.Sleep(1 * time.Second)
	// }
	var i int 
	for {
		if i > 3 {
			break
		}
		fmt.Println(i)
		i += 1
		time.Sleep(300 * time.Millisecond)
	}

	loop:
		for i := 0; i < 10; i++ {
			switch i {
			case 2:
				continue
			case 3:
				continue
			case 8:
				break loop
			default:
				fmt.Printf("%v ", i)
			}
		}
		fmt.Printf("\n")

		items := []item{
			{price: 100.},
			{price: 200.},
			{price: 300.},
		}
		for i := range items {
			items[i].price *= 1.1
		}
		fmt.Printf("items: %+v\n", items)
}

type item struct {
	price float32
}
