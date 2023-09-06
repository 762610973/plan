package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Format(time.DateTime))
	ch := make(chan int)
	defer close(ch)
	timer := time.NewTimer(time.Second * 5)
	go func() {
		select {
		case <-timer.C:
			ch <- 23
		}
	}()
	// 此处会一直阻塞
	data := <-ch
	fmt.Println(data)
	fmt.Println(time.Now().Format(time.DateTime))
}
