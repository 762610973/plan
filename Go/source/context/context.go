package main

import (
	"context"
	"fmt"
)

func main() {
	contextWithValue()
}

func contextWithValue() {
	ctx := context.Background() // parent context
	ctx1 := context.WithValue(ctx, "key1", "val1")
	fmt.Println("get value from ctx1, this value was set by ctx1, value: ", ctx1.Value("key1"))
	ctx2 := context.WithValue(ctx1, "key2", "val2")
	fmt.Println("get value from ctx2, this value was set by ctx2, value: ", ctx2.Value("key2"))
	fmt.Println("get value from ctx2, this value was set by ctx1, value: ", ctx2.Value("key1"))
}
