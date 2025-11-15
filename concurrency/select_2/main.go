package main

import (
	"sync"
	"time"
	"fmt"
)

const bufSize = 3

func main() {
	var wg sync.WaitGroup
	ch := make(chan string, bufSize)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < bufSize; i++ {
			time.Sleep(1000 * time.Millisecond)
			ch <- "hello"
		}
	}()


	for i := 0; i < 3; i++ {
		select {
		case v := <-ch:
			fmt.Println(v)
		default:
			fmt.Println("no msg arrived")
		}
		time.Sleep(1500 * time.Millisecond)
	}
	wg.Wait()


}