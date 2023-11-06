# `sync.WaitGroup`
> 等所有的Add方法调用之后再调用Wait!
## 结构
```go
type WaitGroup struct {
	// 辅助vet工具检查防止复制, 不应该用于值传递.
	noCopy noCopy
	// 高32bit表示计数器, 低32bit表示等待的goroutine数量.
	state atomic.Uint64 // high 32 bits are counter, low 32 bits are waiter count.
	sema  uint32
}
```
## 方法
### `wg.Add()`
```go
// 将delta加到state的高32bit上
// 如果counter为0, 并且waiter大于0, 表示所有被等待的goroutine都完成了
func (wg *WaitGroup) Add(delta int) {
	if race.Enabled {
		if delta < 0 {
			// Synchronize decrements with Wait.
			race.ReleaseMerge(unsafe.Pointer(wg))
		}
		race.Disable()
		defer race.Enable()
	}
	// delta加到高32bit上
	state := wg.state.Add(uint64(delta) << 32)
	// 算数右移32bit, 左边补32个0, 然后类型强转得到counter
	counter := int32(state >> 32)
	// 类型强转, 只取右边32位得到waiter, waiter表示阻塞在Wait()上的goroutine数量.
	w := uint32(state)
	if race.Enabled && delta > 0 && counter == int32(delta) {
		race.Read(unsafe.Pointer(&wg.sema))
	}
	// counter小于0, 直接panic(加上delta之后waiter不能为负数 最小是0)
	if counter < 0 {	
		panic("sync: negative WaitGroup counter")
	}
	// counter > 0
	
	// 正常情况下先调用Add再调用Wait, 此时w是0, counter大与0
	if w != 0 && delta > 0 && counter == int32(delta) {
		// waiter不为0, 说明有调用Wait阻塞的, delta大于0说明在Add, 然后counter==delta说明在初始化添加, 此时同时调用Wait会阻塞住的.
		// Add和Wait不能并发调用.
		panic("sync: WaitGroup misuse: Add called concurrently with Wait")
	}
	// 计数器大于0, waiter为0, 说明没有在wait的goroutine, 还没有到唤醒的时候, 可能都在运行.
	if counter > 0 || w == 0 {
		return
	}
	// 计数器已经为0了, 现在不能同时发生状态突变
		// Add不能与Wait同时发生
		// 如果Wait看到计数器为0, 不会增加Waiter
	// 做一个廉价的判断检测是否滥用.
	if wg.state.Load() != state {	// Add和Wait不能并发调用.
		panic("sync: WaitGroup misuse: Add called concurrently with Wait")
	}
	// 如果计数器为0, waiter不为0, 那么state的值就是waiter的值.
	wg.state.Store(0) // counter==0, 说明所有任务都完成了, 唤醒等待的goroutine.
	for ; w != 0; w-- {
		// 释放信号量, 调用Wait的地方解除阻塞.
		runtime_Semrelease(&wg.sema, false, 0)
	}
}
```
### `wg.Wait()`
> 不断检测state的值, 如果counter为0, 说明所有任务已完成, 调用者不必等待, 直接返回
> 如果counter大于0, 还有任务没完成, 调用者变成了waiter, 需要加入waiter队列, 并且阻塞住自己.
```go
func (wg *WaitGroup) Wait() {
	if race.Enabled {
		race.Disable()
	}
	for {
		state := wg.state.Load()
		// 算数右移, 左侧填32个0, 然后类型强转取右边32位, 得到counter
		counter := int32(state >> 32)
		// 得到waiter
		w := uint32(state)
		if counter == 0 {
			// counter为0, 不需要等待, 直接return.
			if race.Enabled {
				race.Enable()
				race.Acquire(unsafe.Pointer(wg))
			}
			return
		}
		// 调用一次Wait, waiter数量会+1
		if wg.state.CompareAndSwap(state, state+1) {
			if race.Enabled && w == 0 {
				// 等待必须与第一个 Add 同步，需要对此进行建模，以与 Add 中的读取竞争。因此，只能为第一个服务员执行写入，否则并发等待将相互竞争。
				// wait应与第一次调用的Add同步. 只能为第一个waiter服务, 否则会并发等待相互竞争.
				race.Write(unsafe.Pointer(&wg.sema))
			}
			// 阻塞休眠等待
			runtime_Semacquire(&wg.sema)
			// 被唤醒, 不再阻塞, return.
			if wg.state.Load() != 0 {
				panic("sync: WaitGroup is reused before precounterious Wait has returned")
			}
			if race.Enabled {
				race.Enable()
				race.Acquire(unsafe.Pointer(wg))
			}
			return
		}
		// 状态没有修改成功, 开始下一次尝试.
	}
}
```