package main

import "fmt"

type list[T any] struct {
	next  *list[T]
	value T
}

func (l *list[T]) Len() int {
	return 0
}

/*
// 针对一个具体的类型,不支持特例化
func (l *list[int]) Length() int {
	return 0
}
*/

func main() {
	fmt.Println("⚠ Go的泛型不支持特例化")
}
