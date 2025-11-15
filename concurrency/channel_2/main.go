package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch1 := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(<-ch1)
	}()
	ch1 <- 10
	close(ch1)
	v, ok := <-ch1
	fmt.Printf("v: %v, ok: %v\n", v, ok)
	wg.Wait()

	ch2 := make(chan int, 2)
	ch2 <- 1
	ch2 <- 2
	close(ch2)
	v, ok = <-ch2
	fmt.Printf("v: %v, ok: %v\n", v, ok)
	v, ok = <-ch2
	fmt.Printf("v: %v, ok: %v\n", v, ok)
	v, ok = <-ch2
	fmt.Printf("v: %v, ok: %v\n", v, ok)

	ch3 := generateCountStream()
	for v := range ch3 { // ここでチャネルからデータを受信することで、チャネルがcloseされるまでループする
		fmt.Println(v)
	}

	nCh := make(chan struct{})
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("goroutine %v started\n", i)
			<-nCh
			fmt.Println(i)
		}(i)
	}
	time.Sleep(2 * time.Second)
	close(nCh)
	fmt.Println("unblockied by manual close")

	wg.Wait()
	fmt.Println("all goroutines finished")

	ch4 := generateCountStream2()
	for v := range ch4 {
		fmt.Println(v)
		time.Sleep(2 * time.Second)
	}
}

func generateCountStream() <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i <= 5; i++ {
			ch <- i
		}
	}()
	return ch
}

func generateCountStream2() <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i <= 5; i++ {
			ch <- i
			fmt.Println("write")
		}
	}()
	return ch
}