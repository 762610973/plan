package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"
)

func main() {
	// contextWithTimeout()
	// contextWithCancel()
	// contextWithDeadLine()
	// contextWithCancelCase()
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
	exitCh := make(chan struct{})
	fmt.Println(time.Now().Format(time.DateTime))
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
			exitCh <- struct{}{}
		}
	}()
	for {
		select {
		case val := <-ticker.C:
			fmt.Println("ticker: ", val.Format(time.DateTime))
		case <-exitCh:
			fmt.Println(time.Now().Format(time.DateTime))
			return
		}
	}
}

func contextWithCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ticker := time.NewTicker(time.Millisecond * 500)
	sp := 1
	for {
		select {
		case v := <-ticker.C:
			if sp == 5 {
				cancel()
			}
			slog.Info(v.Format(time.DateTime))
			sp++
		case <-ctx.Done():
			slog.Info(ctx.Err().Error())
			// fmt.Println(ctx.Deadline()) withCancel Deadline()没有意义
			return
		}
	}
}

func contextWithDeadLine() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))
	defer cancel()
	select {
	case <-time.After(time.Second * 4):
		cancel()
		fmt.Println(ctx.Err())
		// 两秒后接收到信号
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func contextWithCancelCase() {
	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(nil)
	ticker := time.NewTicker(time.Millisecond * 500)
	sp := 1
	for {
		select {
		case v := <-ticker.C:
			if sp == 5 {
				cancel(errors.New("my error"))
			}
			slog.Info(v.Format(time.DateTime))
			sp++
		case <-ctx.Done():
			slog.Info(ctx.Err().Error())
			// fmt.Println(ctx.Deadline()) CancelCause Deadline()没有意义
			fmt.Println(context.Cause(ctx))
			return
		}
	}
}
