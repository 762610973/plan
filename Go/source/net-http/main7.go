package main

import (
	"fmt"
	"io"
	"net/http"
	"runtime"
	"time"
)

func main() {
	fmt.Println("🤓: 最佳实践是所有client都共用一个transport")
	fmt.Println("🤓: 每个http.Transport内都会维护一个自己的空闲连接池")
	main1()
	//main2()
}
func main1() {
	n := 5
	for i := 0; i < n; i++ {
		httpClient := &http.Client{}
		resp, _ := httpClient.Get("https://www.baidu.com")
		_, _ = io.ReadAll(resp.Body)
		_ = resp.Body.Close()
	}
	time.Sleep(time.Second * 5)
	//  * 三个goroutine, 有一个被复用的TCP连接
	fmt.Printf("goroutine num is %d\n", runtime.NumGoroutine())
}

func main2() {
	n := 5
	for i := 0; i < n; i++ {
		httpClient := &http.Client{
			Transport: &http.Transport{},
		}
		resp, _ := httpClient.Get("https://www.baidu.com")
		_, _ = io.ReadAll(resp.Body)
		_ = resp.Body.Close()
	}
	time.Sleep(time.Second * 1)
	// * 11个goroutine,main goroutine,5个write,5个read
	// * 每个client都创建 一个新的http.Transport,导致底层TCP连接无法复用
	fmt.Printf("goroutine num is %d\n", runtime.NumGoroutine())
}
