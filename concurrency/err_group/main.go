package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 800*time.Millisecond)
	defer cancel()

	// eg, ctx := errgroup.WithContext(context.Background())
	// s := []string{"task1", "fake1", "task2", "fake2", "task3"}
	eg, ctx := errgroup.WithContext(ctx)
	s := []string{"task1", "task2", "task3", "task4"}

	for _, v := range s {
		task := v 
		eg.Go(func() error {
			return doTaskWithContext(ctx, task)
		})
	}
	if err := eg.Wait(); err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Println("finished")
}

func doTask(task string) error {
	if task == "fake1" || task == "fake2" {
		return fmt.Errorf("%v failed", task)
	}
	fmt.Printf("task %v completed\n", task)
	return nil
}

func doTaskWithContext(ctx context.Context, task string) error {
	var t *time.Ticker
	switch task {
	case "fake1":
		t = time.NewTicker(500 * time.Millisecond)
	case "task1":
		t = time.NewTicker(500 * time.Millisecond)
	case "fake2":
		t = time.NewTicker(700 * time.Millisecond)
	case "task2":
		t = time.NewTicker(700 * time.Millisecond)
	default:
		t = time.NewTicker(1000 * time.Millisecond)
	}

	select {
	case <-ctx.Done():
		fmt.Printf("%v cancelled : %v\n", task, ctx.Err())
		return ctx.Err()
	case <-t.C:
		t.Stop()
		if task == "fake1" || task == "fake2" {
			return fmt.Errorf("%v process failed", task)
		}
		fmt.Printf("task %v completed\n", task)
	}

	return nil
}