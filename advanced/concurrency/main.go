// Package main demonstrates advanced Go concurrency patterns:
// context propagation, bounded worker pools, the select multiplexer,
// and channel-based pipelines.
package main

import (
	"context"
	"fmt"
	"time"
)

// runPipeline wires three independent stages together with channels:
// generate -> square -> print. Each stage runs in its own goroutine and
// streams values downstream as soon as they're ready.
func generate(ctx context.Context, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

func square(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

// worker is one member of a bounded worker pool. Workers pull jobs from a
// shared, buffered channel until it's closed, capping concurrency at the
// number of workers started regardless of how many jobs exist.
func worker(ctx context.Context, id int, jobs <-chan int, results chan<- int) {
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				return
			}
			results <- job * 10
		case <-ctx.Done():
			// Honor cancellation so the worker doesn't leak.
			return
		}
	}
}

func runWorkerPool(ctx context.Context, jobCount, workerCount int) []int {
	jobs := make(chan int, jobCount)
	results := make(chan int, jobCount)

	for w := 1; w <= workerCount; w++ {
		go worker(ctx, w, jobs, results)
	}

	for j := 1; j <= jobCount; j++ {
		jobs <- j
	}
	close(jobs)

	collected := make([]int, 0, jobCount)
	for i := 0; i < jobCount; i++ {
		collected = append(collected, <-results)
	}
	return collected
}

// nonBlockingCheck shows the select multiplexer with a default case: it
// never blocks, it just reports whether data was ready right now.
func nonBlockingCheck(ch <-chan int) {
	select {
	case v := <-ch:
		fmt.Println("received:", v)
	default:
		fmt.Println("no value ready, moving on")
	}
}

func main() {
	// Context propagation: a deadline that cancels everything downstream.
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Pipeline pattern.
	nums := generate(ctx, 1, 2, 3, 4, 5)
	squares := square(ctx, nums)
	for v := range squares {
		fmt.Println("pipeline square:", v)
	}

	// Worker pool pattern: 3 workers bound concurrency for 9 jobs.
	results := runWorkerPool(ctx, 9, 3)
	fmt.Println("worker pool results:", results)

	// select multiplexer with a default (non-blocking) case.
	empty := make(chan int)
	nonBlockingCheck(empty)

	// select multiplexing a timeout against context cancellation.
	select {
	case <-time.After(50 * time.Millisecond):
		fmt.Println("timer fired first")
	case <-ctx.Done():
		fmt.Println("context cancelled first:", ctx.Err())
	}
}
