package main

import (
	"fmt"
	"io"
	"net/http"
	"runtime"
)

func main() {
	fmt.Printf("goroutine num is %d\n", runtime.NumGoroutine())
	n := 5
	for i := 0; i < n; i++ {
		resp, _ := http.Get("https://www.baidu.com")
		_, _ = io.ReadAll(resp.Body)
		_ = resp.Body.Close()
	}
	// * 此时会输出3,是因为长连接被推入到连接池了,连接会重新复用
	fmt.Printf("goroutine num is %d\n", runtime.NumGoroutine())
}
