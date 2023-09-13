package main

import "fmt"

/*
* 工厂角色:简单工厂模式的核心,负责实现创建所有实例的内部逻辑,工厂类可以直接被外部调用,创建所需的产品对象
* 抽象产品角色:简单工厂模式所创建的所有对象的父类,负责描述所有实例所共有的接口
* 具体产品角色: 简单工厂模式所创建的 具体实例对象
- 缺点是工厂类职责过重,违背开闭原则, 添加新产品需要修改工厂逻辑, 工厂越来越复杂
*/

var (
	_ fruit = (*apple)(nil)
	_ fruit = (*pear)(nil)
)

// !抽象层
type fruit interface {
	show()
}

// !实现层
type apple struct {
}

func (a *apple) show() {
	fmt.Println("this is apple")
}

type pear struct {
}

func (p *pear) show() {
	fmt.Println("this is pear")
}

// !工厂模块
type factory struct{}

func (f *factory) createFruit(name string) fruit {
	var fr fruit
	switch name {
	case "apple":
		fr = new(apple)
	case "pear":
		fr = new(pear)
	default:
		fr = new(apple)
	}

	return fr
}

func main() {
	f := new(factory)
	a := f.createFruit("apple")
	a.show()
	p := f.createFruit("pear")
	p.show()
}
