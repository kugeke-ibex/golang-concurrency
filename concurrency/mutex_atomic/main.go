package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var wg1 sync.WaitGroup
	var mu sync.Mutex
	var i int
	wg1.Add(2)
	go func() {
		defer wg1.Done()
		mu.Lock()
		defer mu.Unlock()
		i = 1
	}()
	go func() {
		defer wg1.Done()
		mu.Lock()
		defer mu.Unlock()
		i = 2
	}()
	wg1.Wait()
	fmt.Println(i)
	// go run -race main.go: raceが検知された場合は、競合があることを意味する
	// 1と2が競合しているため、結果は不定であり、この事を競合状態(race condition)という

	var wg2 sync.WaitGroup
	var rwmu sync.RWMutex
	var c int

	wg2.Add(4)
	go write(&rwmu, &wg2, &c)
	go read(&rwmu, &wg2, &c)
	go read(&rwmu, &wg2, &c)
	go read(&rwmu, &wg2, &c)
	wg2.Wait()
	fmt.Printf("c1: %d\n", c)
	fmt.Println("finished")

	var wg3 sync.WaitGroup
	var c2 int64

	for i := 0; i < 5; i++ {
		wg3.Add(1)
		go func() {
			defer wg3.Done()
			for j := 0; j < 10; j++ {
				atomic.AddInt64(&c2, 1) // 整数の加算を排他制御で行う
			}
		}()
	}
	wg3.Wait()
	fmt.Printf("c2: %d\n", c2)
	fmt.Println("all goroutines finished")
}

func read(mu *sync.RWMutex, wg *sync.WaitGroup, c *int) {
	defer wg.Done()
	time.Sleep(100 * time.Millisecond)
	mu.RLock() // 読み込みロック: 複数の同時読み込みが可能であるが、書き込みは排他的に行われる
	defer mu.RUnlock()
	fmt.Println("read lock")
	fmt.Println(*c)
	time.Sleep(1 * time.Second)
	fmt.Println("read unlock")
}

func write(mu *sync.RWMutex, wg *sync.WaitGroup, c *int) {
	defer wg.Done()
	mu.Lock()
	defer mu.Unlock()
	fmt.Println("write lock")
	*c++
	time.Sleep(1 * time.Second)
	fmt.Println("write unlock")
}