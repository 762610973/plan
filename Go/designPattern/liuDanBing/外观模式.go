package main

import "fmt"

// !外观模式
type subSystemA struct{}

func (s subSystemA) methodA() {
	fmt.Println("子系统A方法A")
}

type subSystemB struct{}

func (s subSystemB) methodB() {
	fmt.Println("子系统B方法B")
}

type subSystemC struct{}

func (s subSystemB) methodC() {
	fmt.Println("子系统C方法C")
}

// !外观模式, 提供了一个外观类, 简化成一个简单的接口供使用(比如main函数, 通常代码比较少)
// * 对客户端屏蔽子系统组件, 减少了客户端所需处理的对象数目
// * 如果设计不当, 增加的子系统可能需要修改外观类的源代码,违背了开闭原则
type facade struct {
	a subSystemA
	b subSystemB
	c subSystemC
}

func (f facade) methodOne() {
	f.a.methodA()
}

func (f facade) methodTwo() {
	f.a.methodA()
	f.b.methodB()
}

func main() {
	a := subSystemA{}
	b := subSystemB{}
	a.methodA()
	b.methodC()

	f := facade{
		a: subSystemA{},
		b: subSystemB{},
	}
	f.methodOne()
}
