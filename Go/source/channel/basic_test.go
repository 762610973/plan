package main

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

// 向一个已经关闭的channel发送数据会panic: send on closed channel
func Test_channel1(t *testing.T) {
	ch := make(chan struct{})
	close(ch)
	ch <- struct{}{}
}

// 关闭一个已经关闭的channel会panic: close of closed channel
func Test_channel2(t *testing.T) {
	ch := make(chan struct{})
	close(ch)
	close(ch)
}

// 从一个未关闭且没有数据的channel中读取数据, 会死锁fatal error: all goroutines are asleep - deadlock!
func Test_channel3(t *testing.T) {
	ch := make(chan struct{}, 1)
	ch <- struct{}{}
	val, ok := <-ch
	fmt.Println(val, ok)
	fmt.Println(<-ch)
}

// 向一个已经满的channel发送数据, 会死锁, 因为没有接收方
func Test_channel4(t *testing.T) {
	ch := make(chan struct{}, 1)
	ch <- struct{}{}
	ch <- struct{}{}
}

// 从一个已经关闭的channel中读取数据, 如果已经读完, 再读的数据就是对应类型的零值
// 即可以从一个已经关闭的channel中持续读
func Test_channel5(t *testing.T) {
	ch := make(chan int, 1)
	ch <- rand.Int()
	close(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// 不要使用for range channel, 而是使用一个值去接收
func Test_channel6(t *testing.T) {
	ch := make(chan int, 4)
	for i := 0; i < 4; i++ {
		ch <- 3
	}
	close(ch)
	for range ch {
		fmt.Println(<-ch)
	}
}

// 使用cloe channel之后, for v := range channel, 可以读完所有的值, 而不会读到零值
func Test_channel7(t *testing.T) {
	ch := make(chan int, 4)
	for i := 0; i < 4; i++ {
		ch <- rand.Intn(5)
	}
	close(ch)
	for v := range ch {
		fmt.Println(v)
	}
}

// 优雅关闭channel
func Test_channel8(t *testing.T) {
	var num int
	ch := make(chan int, 5)
	go func() {
		for {
			if num == 20 {
				// 生产者关闭
				close(ch)
				break
			}
			ch <- rand.Intn(10)
			num++
		}
	}()
	for {
		// close之后, 读取的ok就是false了
		val, ok := <-ch
		if ok {
			fmt.Println(val)
		} else {
			fmt.Println("the channel is closed")
			return
		}
	}
}

// 从channel中读取数据
func Test_channel9(t *testing.T) {
	var num int
	ch := make(chan int, 5)
	go func() {
		for {
			if num == 20 {
				// 生产者关闭
				close(ch)
				break
			}
			ch <- rand.Intn(10)
			num++
		}
	}()

	for v := range ch {
		fmt.Println(v)
	}
}

// 关闭ticker
func Test_ticker(t *testing.T) {
	ticker := time.NewTicker(time.Second)
	exit := make(chan struct{})
	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
		exit <- struct{}{}
	}()
	for {
		select {
		case v := <-ticker.C:
			fmt.Println(v)
		case <-exit:
			return
		}
	}
}

// 无缓冲channel读取
func Test_channel10(t *testing.T) {
	var num int
	ch := make(chan int)
	go func() {
		for {
			if num == 20 {
				// 生产者关闭
				close(ch)
				break
			}
			ch <- rand.Intn(10)
			num++
		}
	}()
	for v := range ch {
		time.Sleep(time.Second)
		fmt.Println(v)
	}
}

// 多生产者单消费者, 同时阻塞式消费
func Test_channel11(t *testing.T) {
	ch := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 2; i++ {
			time.Sleep(time.Second)
			ch <- "first goroutine"
		}
	}()
	go func() {
		wg.Done()
		for i := 0; i < 3; i++ {
			time.Sleep(time.Millisecond * 500)
			ch <- "second goroutine"
		}
	}()
	go func() {
		wg.Wait()
		close(ch)
	}()

	for i := range ch {
		fmt.Println(i)
	}
}
