package main

import (
	"fmt"
	"sync"
)

var num = 5

type empty struct{}

func main() {
	mod1()
	mod2()
}

func mod1() {
	exitCh := make(chan bool)   //控制结束退出的信号
	chStart := make(chan empty) //chan通信 控制起始同步阻塞的
	chEnd := make(chan empty)   //chan通信 控制结尾同步阻塞的
	go func() {
		for i := 0; i <= num; i++ {
			//生产 控制起始
			chStart <- empty{}
			fmt.Println("A ", i)
			//暂停 等待结尾信号
			<-chEnd
		}
		//发送结束之后退出
		exitCh <- true
	}()

	go func() {
		for i := 0; i <= num; i++ {
			//消费 结束起始
			<-chStart
			fmt.Println("B ", i)
			//发送结尾信号
			chEnd <- empty{}
		}
	}()
	//阻塞等待退出信号
	<-exitCh
}

func mod2() {
	var wg sync.WaitGroup
	wg.Add(2)
	sigA := make(chan empty, 1)
	defer close(sigA)
	sigB := make(chan empty)
	defer close(sigB)
	// goroutine a print 'a'.
	go func() {
		for i := 0; i < num; i++ {
			<-sigA
			fmt.Printf("goroutine a print 'a', num: %d\n", i)
			sigB <- empty{}
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < num; i++ {
			<-sigB
			fmt.Printf("goroutine b print 'b', num: %d\n", i)
			sigA <- empty{}
		}
		wg.Done()
	}()
	sigA <- empty{}
	wg.Wait()
}
