# `sync.Once`
> 核心思想: 互斥锁+双检测, 避免并发调用初始化资源等问题.
## 结构
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
## 方法
### `once.Do()`
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
### `once.OnceFunc()`
```go
// 入参为一个func(), 返回一个func(), 可以并发调用, 多次调用只会执行一次
func OnceFunc(f func()) func() {
	var (
		once  Once
		// 标识是否发生了panic
		valid bool
		p     any
	)
	// 只需要构建一次内部闭包, 就能减少在fast path上的成本
	g := func() {
		defer func() {
			p = recover()
			if !valid {
				// 立即重新panic, 因此在第一次调用时, 可以获得完整的堆栈跟踪到f
				panic(p)
			}
		}()
		f()
		valid = true // 正常执行设置为true, 如果发生panic, 这一行不会执行, 所以defer中可以检测到发生了panic
	}
	// 返回一个func, 使用once.Do()包装, 调用时会执行g, 在g里面会执行f
	return func() {
		once.Do(g)
		// 如果valid还是false, 说明没有正常执行
		//! 考虑并发问题, 如果f注定panic, 同时有多个goroutine并发执行.
		//! 第二次往后的并发访问不会去执行g, 则会进入到下面的逻辑, 也可以通过panic跟踪到完整堆栈.
		
		if !valid {
			panic(p)
		}
	}
}
```
### `once.OnceValue()`
```go
func OnceValue[T any](f func() T) func() T {
	var (
		once   Once
		valid  bool
		p      any
		result T
	)
	g := func() {
		defer func() {
			p = recover()
			if !valid {
				panic(p)
			}
		}()
		result = f()
		valid = true
	}
	return func() T {
		once.Do(g)
		if !valid {
			panic(p)
		}
		return result
	}
}
```
### `once.OnceValues()`
```go
func OnceValues[T1, T2 any](f func() (T1, T2)) func() (T1, T2) {
	var (
		once  Once
		valid bool
		p     any
		r1    T1
		r2    T2
	)
	g := func() {
		defer func() {
			p = recover()
			if !valid {
				panic(p)
			}
		}()
		r1, r2 = f()
		valid = true
	}
	return func() (T1, T2) {
		once.Do(g)
		if !valid {
			panic(p)
		}
		return r1, r2
	}
}
```