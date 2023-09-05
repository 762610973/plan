package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("对于无缓冲channel, 对于channel的操作需要先写一个goroutine, 否则会一直阻塞在那里, 造成死锁")
	order1()
	order2()
}

func order1() {
	fmt.Println("order1")
	wg := sync.WaitGroup{}
	ch := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(<-ch)
	}()
	ch <- 1
	wg.Wait()
	fmt.Println("exit")
}

func order2() {
	fmt.Println("order2")
	wg := sync.WaitGroup{}
	ch := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		ch <- 1
	}()
	fmt.Println(<-ch)
	wg.Wait()
	fmt.Println("exit")
}
