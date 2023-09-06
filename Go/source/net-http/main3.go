package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func main() {
	fmt.Printf("goroutine num is %d\n", runtime.NumGoroutine())
	n := 5
	for i := 0; i < n; i++ {
		resp, _ := http.Get("https://www.baidu.com")
		_ = resp.Body.Close()
	}
	time.Sleep(time.Millisecond)
	// *增加休眠之后,只有一个main goroutine
	fmt.Printf("goroutine num is %d\n", runtime.NumGoroutine())
	// 输出3,main goroutine,read goroutine,write goroutine,长连接没有断开,但是不会在下一次复用
	// 长连接释放需要时间
}
