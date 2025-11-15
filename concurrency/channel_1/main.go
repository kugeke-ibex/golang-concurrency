package main

import (
	"fmt"
	"runtime"
	// "sync"
	// "time"
)

func main() {
	// ch := make(chan int)
	// ch <- 10 // チャネルの呼び出し前に先にデータを送信する
	// fmt.Println(<-ch)  // 上記によりdeadlockが発生する

	// ch := make(chan int)
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func () {
	// 	defer wg.Done()
	// 	ch <- 10
	// 	time.Sleep(1 * time.Second)
	// }()
	// fmt.Println(<-ch)
	// wg.Wait()

	ch1 := make(chan int)
	go func() {
		// goroutineが終了するまで待つことを「goroutine leak」と言う
		fmt.Println(<-ch1) // ここでチャネルのデータを受信することで、goroutineが終了する
	}()
	ch1 <- 10 // ここでチャネルのデータを送信する
	fmt.Printf("num of working goroutines: %d\n", runtime.NumGoroutine())

	ch2 := make(chan int, 1) // バッファリングされたチャネルを作成する
	ch2 <- 2 // バッファリングされたチャネルを使うことで、 チャネルにデータを送信できる
	// ch2 <- 3 // バッファリングされたチャネルが満杯のため、deadlockが発生する
	fmt.Println(<-ch2) // バッファリングされたチャネルを使うことで、 チャネルからデータを受信できる	
	// 実行順序を逆にすると、チャネルが空のため、deadlockが発生する
}
