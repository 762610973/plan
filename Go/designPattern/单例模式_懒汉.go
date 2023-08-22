package main

import (
	"sync"
	"sync/atomic"
)

type lazySingle struct{}

var instance *lazySingle

var mu sync.Mutex

// GetLazyInstanceWithMutex 使用互斥锁
func GetLazyInstanceWithMutex() *lazySingle {
	mu.Lock()
	defer mu.Unlock()
	if instance == nil {
		instance = new(lazySingle)
		return instance
	}

	return instance
}

var initialized int32

// GetLazyInstanceWithAtomic 使用原子变量+互斥锁
func GetLazyInstanceWithAtomic() *lazySingle {
	if atomic.LoadInt32(&initialized) == 1 {
		return instance
	}
	// 如果没有, 则加锁申请,防止其它的goroutine也生成一份实例
	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		instance = new(lazySingle)
		// 设置原子变量的标记
		atomic.StoreInt32(&initialized, 1)
	}

	return instance
}

var once sync.Once

func GetLazyInstanceWithOnce() *lazySingle {
	once.Do(func() {
		instance = new(lazySingle)
	})

	return instance
}
