# `sync.Mutex`
## 结构
```go
type Mutex struct {
	// state是一个复合型字段, 一个字段包含多个意义, 可以通过尽可能少的内存来实现互斥锁
	// mutexWaiters|mutexStarving|mutexWoken|mutexLocked
	state int32
	sema  uint32
}
const (
	mutexLocked = 1 << iota // mutex is locked
	// 当前是否已经存在被唤醒的goroutine, 因为mutex需要知道当前是否已经存在试图获取锁的goroutine
	// 存在的话就不用再额外去唤醒锁了, 因为唤醒了也是去竞争, 不如让运行中的goroutine直接拿到
	mutexWoken				// 2
	mutexStarving			// 4
	mutexWaiterShift = iota // 3
	starvationThresholdNs = 1e6
)
```
## 方法
```go
func (m *Mutex) Lock() {
	// Fast path: 幸运case, 直接获取到锁.
	// cas操作, 如果old是0, 说明四个标记位都为0没有在使用锁的, cas可以直接获取到锁, 将state设置为1, 即最后一位mutexLocked为1, 表示已经被占有了
	if atomic.CompareAndSwapInt32(&m.state, 0, mutexLocked) {
		if race.Enabled {
			race.Acquire(unsafe.Pointer(m))
		}
		return
	}
	m.lockSlow()
}
func (m *Mutex) lockSlow() {
	var waitStartTime int64	// goroutine的等待时间
	starving := false	// 此goroutine的饥饿标记
	awoke := false		// 唤醒模式标记
	iter := 0			// 自旋次数
	old := m.state 		// 保存当前锁的状态
	for {
		// 锁是非饥饿状态, 锁还没被释放, 尝试自旋
		// mutexLocked | mutexStarving: 按位或运算
		/*
		state: waiter|starving|woken|locked
		mutexLocked是1(001)
		mutexStarving是4(100)
		mutexLocked|mutexStarving = 101
		old&101 => 的结果取决于old的starving和locked位, 因为只有这俩是1(按位与), 前面的都要补0, 然后locked是1, 就取决于starving了
		如果这个最终结果 == mutexLocked(1), 那么starving只能是0, 表示处于饥饿状态
		! old&(mutexLocked|mutexStarving)作用是判断是否处于饥饿状态, 拿到starving位的值
		! old&mutexWoken作用是判断是否有被唤醒的goroutine, 因为结果取决于old的woken位的值
		10
		*/
		// 如果上了锁并且可以自旋, 还不是饥饿状态
		if old&(mutexLocked|mutexStarving) == mutexLocked && runtime_canSpin(iter) {
			if !awoke && old&mutexWoken == 0 && old>>mutexWaiterShift != 0 &&
				atomic.CompareAndSwapInt32(&m.state, old, old|mutexWoken) {
				awoke = true
			}
			// 自旋等待, 目的是不直接挂起, 有可能锁马上就会释放
			runtime_doSpin()
			iter++
			old = m.state	// 再次获取锁的状态，之后会检查是否锁被释放了
			continue
		}
		new := old
		// 所有goroutine的目的都是获取锁, 但是如果当前处于饥饿状态, 则都不允许, 统统去排队.
		if old&mutexStarving == 0 {
			new |= mutexLocked	// 非饥饿状态, 加锁
		}
		// 如果锁被持有, 或者处于饥饿状态, 那么waiter+1
		if old&(mutexLocked|mutexStarving) != 0 {
			new += 1 << mutexWaiterShift
		}
		// 如果锁被持有, 并且goroutine进入饥饿模式, 那么切换mutex为饥饿状态, 要求mutex按照饥饿模式执行
		if starving && old&mutexLocked != 0 {
			new |= mutexStarving
		}
		if awoke {
			// goroutine已经被唤醒, 是清醒的. 所以需要重置标志位
			if new&mutexWoken == 0 {
				throw("sync: inconsistent mutex state")
			}
			new &^= mutexWoken	// 新状态清除唤醒标志
		}
		if atomic.CompareAndSwapInt32(&m.state, old, new) {
			// 老状态没上锁并且处于正常模式, 那么一定上锁成功
			// 老状态没上锁, 但是是饥饿模式, 需要排队等候
			if old&(mutexLocked|mutexStarving) == 0 {
				break // locked the mutex with CAS
			}
			// 如果是等待过的goroutine, 则进入队列头部, 之后会优先出队
			queueLifo := waitStartTime != 0
			if waitStartTime == 0 {
				waitStartTime = runtime_nanotime()
			}
			// 休眠
			runtime_SemacquireMutex(&m.sema, queueLifo, 1)
			// 被唤醒, 先判断是否应该进入饥饿状态
			starving = starving || runtime_nanotime()-waitStartTime > starvationThresholdNs
			old = m.state
			// 如果锁处于饥饿模式,那么被唤醒的协程是一定能拿到锁的.
			// 注意上面的饥饿状态,协程进入饥饿状态不代表锁进入饥饿模式,需要这个饥饿状态的锁把锁设置成饥饿模式才行.
			// 这里实际是一个手递手的过程
			if old&mutexStarving != 0 {
				// 不可能的状态预警
				if old&(mutexLocked|mutexWoken) != 0 || old>>mutexWaiterShift == 0 {
					throw("sync: inconsistent mutex state")
				}
				// waiter-1和上锁
				delta := int32(mutexLocked - 1<<mutexWaiterShift)
				if !starving || old>>mutexWaiterShift == 1 {
					// Exit starvation mode.
					// Critical to do it here and consider wait time.
					// Starvation mode is so inefficient, that two goroutines
					// can go lock-step infinitely once they switch mutex
					// to starvation mode.
					// 判断是否需要退出饥饿模式
					delta -= mutexStarving
				}
				atomic.AddInt32(&m.state, delta)
				break
			}
			awoke = true
			iter = 0
		} else {
			old = m.state
		}
	}

	if race.Enabled {
		race.Acquire(unsafe.Pointer(m))
	}
}

func (m *Mutex) Unlock() {
	if race.Enabled {
		_ = m.state
		race.Release(unsafe.Pointer(m))
	}

	// Fast path: drop lock bit.
	new := atomic.AddInt32(&m.state, -mutexLocked)
	if new != 0 {
		// Outlined slow path to allow inlining the fast path.
		// To hide unlockSlow during tracing we skip one extra frame when tracing GoUnblock.
		m.unlockSlow(new)
	}
}

func (m *Mutex) unlockSlow(new int32) {
	if (new+mutexLocked)&mutexLocked == 0 {
		fatal("sync: unlock of unlocked mutex")
	}
	if new&mutexStarving == 0 {
		old := new
		for {
			// If there are no waiters or a goroutine has already
			// been woken or grabbed the lock, no need to wake anyone.
			// In starvation mode ownership is directly handed off from unlocking
			// goroutine to the next waiter. We are not part of this chain,
			// since we did not observe mutexStarving when we unlocked the mutex above.
			// So get off the way.
			if old>>mutexWaiterShift == 0 || old&(mutexLocked|mutexWoken|mutexStarving) != 0 {
				return
			}
			// Grab the right to wake someone.
			new = (old - 1<<mutexWaiterShift) | mutexWoken
			if atomic.CompareAndSwapInt32(&m.state, old, new) {
				runtime_Semrelease(&m.sema, false, 1)
				return
			}
			old = m.state
		}
	} else {
		// Starving mode: handoff mutex ownership to the next waiter, and yield
		// our time slice so that the next waiter can start to run immediately.
		// Note: mutexLocked is not set, the waiter will set it after wakeup.
		// But mutex is still considered locked if mutexStarving is set,
		// so new coming goroutines won't acquire it.
		runtime_Semrelease(&m.sema, true, 1)
	}
}
func (m *Mutex) TryLock() bool {
	old := m.state
	if old&(mutexLocked|mutexStarving) != 0 {
		return false
	}

	// There may be a goroutine waiting for the mutex, but we are
	// running now and can try to grab the mutex before that
	// goroutine wakes up.
	if !atomic.CompareAndSwapInt32(&m.state, old, old|mutexLocked) {
		return false
	}

	if race.Enabled {
		race.Acquire(unsafe.Pointer(m))
	}
	return true
}
```