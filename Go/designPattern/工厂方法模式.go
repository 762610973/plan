package main

import "fmt"

// 简单工厂模式 + 开闭原则 = 工厂方法模式
// 新的类只需要实现fruit接口和abstractFactory接口即可

// !水果类
type fruit interface {
	show()
}

// !工厂类
type abstractFactory interface {
	createFruit() fruit
}

type peach struct{}

func (p *peach) show() {
	fmt.Println("this is peach")
}

type banana struct{}

func (b *banana) show() {
	fmt.Println("this is banana")
}

// !基础的工厂模块

// peach工厂
type peachFactory struct{}

func (p *peachFactory) createFruit() fruit {
	return new(peach)
}

// banana工厂
type bananaFactory struct{}

func (p *bananaFactory) createFruit() fruit {
	return new(banana)
}

func main() {
	// 需要一个具体的peach对象
	// 先创建一个具体的peach工厂
	var peachFac abstractFactory
	peachFac = new(peachFactory)
	// 然后生产一个具体的peach
	var peach1 fruit
	peach1 = peachFac.createFruit()
	// 多态
	peach1.show()

	// 需要一个具体的banana对象
	bananaFac := new(bananaFactory)
	banana1 := bananaFac.createFruit()
	banana1.show()
}
