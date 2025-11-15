package main

import (
	"sync"
	"time"
	"runtime"
	"log"
	"context"
	"fmt"
	"os"
	"runtime/trace"
)

func main () {
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done() // wg.Done() を呼び出すことで、wg.Add(1) で増やしたカウントを減らす
	// 	fmt.Println("goroutine invoked") // main goroutineが先に終了してしまうので、printされない
	// }()
	// wg.Wait() // wg.Done() が呼び出されるまで、main goroutine がブロックされる
	// fmt.Printf("num of working goroutines: %d\n", runtime.NumGoroutine())
	
	f, err := os.Create("trace.out") // trace.out ファイルを作成
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	defer func () {
		if err := f.Close(); err != nil {
			log.Fatalln("Error: ", err)
		}
	}()
	if err := trace.Start(f); err != nil { // trace.out ファイルを開いて、trace.Start でトレースを開始
		log.Fatalln("Error: ", err)
	}
	defer trace.Stop() // trace.out ファイルを閉じて、トレースを停止
	ctx, t := trace.NewTask(context.Background(), "main") // main goroutine のトレースを開始
	defer t.End() // main goroutine のトレースを停止
	fmt.Println("The number of logical CPU Cores: ", runtime.NumCPU()) // 論理CPUコア数を表示
	// task(ctx, "Task1")
	// task(ctx, "Task2")
	// task(ctx, "Task3")
	
	var wg sync.WaitGroup
	wg.Add(3)
	go cTask(ctx, &wg, "Task1")
	go cTask(ctx, &wg, "Task2")
	go cTask(ctx, &wg, "Task3")
	wg.Wait()
		
	s := []int{1, 2, 3}
	for _, i := range s {
		wg.Add(1)
		go func () {
			defer wg.Done()
			fmt.Println(i)
			}()
		}
	wg.Wait()
	fmt.Println("main func finish")
}

func task(ctx context.Context, name string) {
	defer trace.StartRegion(ctx, name).End() // name のタスクのトレースを開始
	time.Sleep(1 * time.Second) 
	fmt.Println("task", name)
}

func cTask(ctx context.Context, wg *sync.WaitGroup, name string) {
	defer trace.StartRegion(ctx, name).End() // name のタスクのトレースを開始
	defer wg.Done() // wg.Done() を呼び出すことで、wg.Add(1) で増やしたカウントを減らす
	time.Sleep(1 * time.Second)
	fmt.Println("cTask", name)
}
// go tool trace trace.out でトレースを表示
