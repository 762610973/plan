package main

import (
	"sync"
)

var (
	_ = sync.WaitGroup{}
	_ = sync.Once{}
	_ = sync.OnceFunc(func() {})
	_ = sync.Cond{}
	_ = sync.Pool{}
	_ = sync.Mutex{}
	_ = sync.RWMutex{}
	_ = sync.Map{}
)

func main() {

}
