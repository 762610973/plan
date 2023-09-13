package main

import "fmt"

// !抽象工厂方法模式是针对产品族进行生产产品
// - 针对产品族进行添加,符合开闭原则
// - 针对产品等级结构进行添加,不符合开闭原则

// !抽象层

type abstractApple interface {
	showApple()
}

type abstractOrange interface {
	showOrange()
}

type abstractFactory interface {
	createApple() abstractApple
	createOrange() abstractOrange
}

// ! 实现层

// 中国产品族

type chinaApple struct{}

func (c chinaApple) showApple() {
	fmt.Println("this is china apple")
}

type chinaOrange struct{}

func (o chinaOrange) showOrange() {
	fmt.Println("this is china orange")
}

// 中国工厂
type chinaFactory struct{}

func (c chinaFactory) createApple() abstractApple {
	var a abstractApple
	a = new(chinaApple)

	return a
}

func (c chinaFactory) createOrange() abstractOrange {
	var o abstractOrange
	o = new(chinaOrange)

	return o
}

// 日本产品族

type japanApple struct{}

func (c japanApple) showApple() {
	fmt.Println("this is japan apple")
}

type japanOrange struct{}

func (o japanOrange) showOrange() {
	fmt.Println("this is japan orange")
}

// 日本工厂
type japanFactory struct{}

func (c japanFactory) createApple() abstractApple {
	var a abstractApple
	a = new(japanApple)

	return a
}

func (c japanFactory) createOrange() abstractOrange {
	var o abstractOrange
	o = new(japanOrange)

	return o
}

func main() {
	//  需要中国的苹果,橙子
	//  创建工厂
	var chinaFac abstractFactory
	chinaFac = new(chinaFactory)
	// - 生成苹果
	var cApple abstractApple
	cApple = chinaFac.createApple()
	// 调用方法
	cApple.showApple()

	// - 生成橙子
	var cOrange abstractOrange
	cOrange = chinaFac.createOrange()
	cOrange.showOrange()
}
