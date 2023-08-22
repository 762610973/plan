package main

import "fmt"

// ! 保证一个类永远只能有一个对象, 且该对象的功能依然能被其他模块使用
// - 饿汉式单例模式
// single 保证私有
type single struct{}

// instance 如果是对外暴露的,外部可能会更改这个指针,所以必须是私有的
var instance *single = new(single)

// GetInstance 对外提供读方法
func GetInstance() *single {
	return instance
}

func (s *single) DoSomething() {
	fmt.Println("单例模式")
}

func main() {
	s1 := GetInstance()
	s1.DoSomething()
	s2 := GetInstance()
	if s1 == s2 {
		fmt.Println("s1 == s2")
	}
}
