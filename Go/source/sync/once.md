# `sync.Once`
> 核心思想: 互斥锁+双检测, 避免并发调用初始化资源等问题.
### 结构
```go
type Once struct {
	// done在结构体第一位, 因为它是一个hot path, 会被经常使用, 放在第一位在某些机器指令更紧凑, 一些机器上减少指令
	// fast path的一个好处是此方法可以内联.
	done uint32
	// 如果不适用mutex, 可能有多个goroutine并发访问, 如果第一次执行的fn很慢, 后续调用的goroutine已经看到执行过了
	// 但是获取初始化资源的时候, 可能得到空值或者nil, 因为fn还没有执行完
	// 互斥锁机制只有一个goroutine进行初始化
	m    Mutex
}
```
### 方法
```go
func (o *Once) Do(f func()) {
	// 原子操作加载done, 如果为0, 执行doSlow方法
	if atomic.LoadUint32(&o.done) == 0 {
		// 允许fast path的内联.
		// 如果有多个goroutine并发调用, 就会进入doSlow方法.
		o.doSlow(f)
	}
}

func (o *Once) doSlow(f func()) {
	// 加锁之后, 并发的goroutine会等待f完成, 而不会多次执行f
	o.m.Lock()
	defer o.m.Unlock()
	// 再次判断是否为0, 如果为0, 是第一次执行, 将done设置为1
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}
```