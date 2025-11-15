package main

import (
	"context"
	"sync"
	"time"
	"fmt"
)

func main() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), 1500 * time.Millisecond)
	defer cancel()

	wg.Add(2)
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		ch1 <- "A"
	}()
	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		ch2 <- "B"
	}()

	loop:
		for ch1 != nil || ch2 != nil {
			select {
			case <-ctx.Done():
				fmt.Println("timed out")
				break loop
			case v := <-ch1:
				fmt.Println(v)
			case v := <-ch2:
				fmt.Println(v)
			}
		}
	wg.Wait()
	fmt.Println("all goroutines finished")
}