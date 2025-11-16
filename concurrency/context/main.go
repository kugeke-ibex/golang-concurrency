package main

import (
	"context"
	"sync"
	"time"
	"fmt"
)

func main() {
	var wg1 sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), 600*time.Millisecond)
	defer cancel()

	wg1.Add(3)
	go subTask(ctx, &wg1, "a")
	go subTask(ctx, &wg1, "b")
	go subTask(ctx, &wg1, "c")
	wg1.Wait()

	var wg2 sync.WaitGroup
	ctx2, cancel2 := context.WithCancel(context.Background())
	defer cancel2()
	wg2.Add(1)
	go func() {
		defer wg2.Done()
		v, err := criticalTask(ctx2)
		if err != nil {
			fmt.Printf("critical task cancelled due to: %v\n", err)
			cancel2()

			return
		}
		fmt.Println("success", v)
	}()

	wg2.Add(1)
	go func() {
		defer wg2.Done()
		v, err := normalTask(ctx2)
		if err != nil {
			fmt.Printf("normal task cancelled due to: %v\n", err)
			return 
		}
		fmt.Println("success", v)
	}()
	wg2.Wait()


	ctx3, cancel3 := context.WithDeadline(context.Background(), time.Now().Add(20*time.Millisecond))
	defer cancel3()
	ch := subTaskWithDeadline(ctx3)
	v, ok := <-ch
	if ok {
		fmt.Println(v)
	}
	fmt.Println("finished")
}

func subTask(ctx context.Context, wg *sync.WaitGroup, id string) {
	defer wg.Done()
	t := time.NewTimer(500 * time.Millisecond)

	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		case <-t.C:
			t.Stop()
			fmt.Printf("subTask %s done\n", id)
			return
		}
	}
}

func subTaskWithDeadline(ctx context.Context) <-chan string{
	ch := make(chan string)
	go func() {
		defer close(ch)
		deadline, ok := ctx.Deadline()
		if ok {
			if deadline.Sub(time.Now().Add(30*time.Millisecond)) < 0 {
				fmt.Println("impossible to meet deadline")
				return
			}
		}
		time.Sleep(30 * time.Millisecond)
		ch <- "success"	
	}()
	return ch
}

func criticalTask(ctx context.Context) (string, error){
	ctx, cancel := context.WithTimeout(ctx, 800*time.Millisecond)
	defer cancel()
	
	t := time.NewTimer(1000 * time.Millisecond)
	select {
		case <-ctx.Done():
			return "", ctx.Err()
		case <-t.C:
			t.Stop()
	}
	return "A", nil
}

func normalTask(ctx context.Context) (string, error){	
	t := time.NewTimer(3000 * time.Millisecond)
	select {
		case <-ctx.Done():
			return "", ctx.Err()
		case <-t.C:
			t.Stop()
	}
	return "B", nil
}