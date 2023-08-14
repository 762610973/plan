package main

import (
	"fmt"
	"io"
	"net/http"
	"runtime"
	"time"
)

func main() {
	fmt.Println("🤓: 将请求完之后的休眠时间改为1s")
	fmt.Printf("goroutine num is %d\n", runtime.NumGoroutine())
	tr := &http.Transport{
		MaxIdleConns:    100,
		IdleConnTimeout: 3 * time.Second,
	}

	n := 5
	for i := 0; i < n; i++ {
		req, _ := http.NewRequest("POST", "https://www.baidu.com", nil)
		req.Header.Add("content-type", "application/json")
		client := &http.Client{
			Transport: tr,
			Timeout:   3 * time.Second,
		}
		resp, _ := client.Do(req)
		_, _ = io.ReadAll(resp.Body)
		_ = resp.Body.Close()
	}
	time.Sleep(time.Second * 1)
	// 改为1之后,goroutine数量为3(main goroutine+read,write goroutine),说明有一个网络连接没有断开,这是一个长连接
	fmt.Printf("goroutine num is %d\n", runtime.NumGoroutine())
}
