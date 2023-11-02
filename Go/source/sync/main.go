package main

import (
	"sync"
)

var (
	_ = sync.WaitGroup{}
	_ = sync.Once{}
	_ = sync.OnceFunc(func() {})
)
