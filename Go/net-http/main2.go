package main

import (
	"fmt"
	"io"
	"net/http"
	"runtime"
	"time"
)

func main() {
	fmt.Println("🤓: 将请求完之后的休眠时间改为1s,在header中增加字段,关闭长连接(http1.1默认开启长连接)")
	fmt.Printf("goroutine num is %d\n", runtime.NumGoroutine())
	tr := &http.Transport{
		MaxIdleConns:    100,
		IdleConnTimeout: 3 * time.Second,
	}

	n := 5
	for i := 0; i < n; i++ {
		req, _ := http.NewRequest("POST", "https://www.baidu.com", nil)
		req.Header.Add("content-type", "application/json")
		// header中增加字段,关闭长连接
		req.Header.Add("connection", "close")
		client := &http.Client{
			Transport: tr,
			Timeout:   3 * time.Second,
		}
		resp, _ := client.Do(req)
		_, _ = io.ReadAll(resp.Body)
		_ = resp.Body.Close()
	}
	time.Sleep(time.Second * 1)
	fmt.Printf("goroutine num is %d\n", runtime.NumGoroutine())
}
