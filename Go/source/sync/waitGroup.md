# `sync.WaitGroup`
## 结构
```go
type WaitGroup struct {
	noCopy noCopy
	// 高32bit表示计数器, 低32bit表示等待的goroutine数量.
	state atomic.Uint64 // high 32 bits are counter, low 32 bits are waiter count.
	sema  uint32
}
```