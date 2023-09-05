package main

import (
	"errors"
	"fmt"
)

type Error interface {
	unwrap() error
}

type myError struct {
}

func (m myError) Error() string {
	return "myError"
}

func (m myError) unwrap() error {
	return errors.New("my error")
}

func unwrap(err error) error {
	// 由于要对err解包,想知道err的前世, 所以要将err转换为Error类型,然后调用方法
	u, ok := err.(Error)
	if !ok {
		return nil
	}
	return u.unwrap()
}

func main() {
	m := myError{}
	fmt.Println(unwrap(m))
}
