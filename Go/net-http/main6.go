package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func main() {
	fmt.Println("🤓: 请求后不读也不关闭,会造成goroutine泄露")
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
			// 增加了timeout,不会出现goroutine泄露的问题
			//Timeout:   3 * time.Second,
			Transport: tr,
		}
		resp, _ := client.Do(req)
		_ = resp
	}
	time.Sleep(time.Second * 5)
	fmt.Printf("goroutine num is %d\n", runtime.NumGoroutine())
}
