package main

import (
	"fmt"
	"io"
	"net/http"
	"runtime"
	"time"
)

func main() {
	fmt.Println("🤓: Go为每个网络句柄创建两个goroutine,一个用于读数据,一个用于写数据")
	fmt.Printf("goroutine num is %d\n", runtime.NumGoroutine())
	tr := &http.Transport{
		MaxIdleConns: 100,
		//是空闲（保持活动状态）连接在关闭自身之前保持空闲的最长时间。零意味着没有限制
		IdleConnTimeout: 3 * time.Second,
	}
	n := 5
	for i := 0; i < n; i++ {
		req, _ := http.NewRequest("POST", "https://www.baidu.com", nil)
		req.Header.Add("content-type", "application/json")
		client := &http.Client{
			Transport: tr,
			//指定此客户端发出的请求的时间限制。包括连接时间、任何重定向和读取响应正文
			Timeout: 3 * time.Second,
		}
		resp, _ := client.Do(req)
		_, _ = io.ReadAll(resp.Body)
		_ = resp.Body.Close()
	}
	time.Sleep(time.Second * 5)
	fmt.Printf("goroutine num is %d\n", runtime.NumGoroutine())
}
