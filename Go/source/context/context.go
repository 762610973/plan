package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"
)

func main() {
	contextWithValue()
	contextWithTimeout()
}

func contextWithValue() {
	ctx := context.Background() // parent context
	ctx1 := context.WithValue(ctx, "key1", "val1")
	fmt.Println("get value from ctx1, this value was set by ctx1, value: ", ctx1.Value("key1"))
	ctx2 := context.WithValue(ctx1, "key2", "val2")
	fmt.Println("get value from ctx2, this value was set by ctx2, value: ", ctx2.Value("key2"))
	fmt.Println("get value from ctx2, this value was set by ctx1, value: ", ctx2.Value("key1"))
}

func contextWithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	ctx.Done()
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Deadline())
			if errors.Is(ctx.Err(), context.DeadlineExceeded) {
				fmt.Println(true)
			}
			fmt.Println(ctx.Err())
			os.Exit(0)
		}
	}()
	for {
		select {
		case val := <-ticker.C:
			fmt.Println("ticker: ", val.Format(time.DateTime))
		}
	}

}
