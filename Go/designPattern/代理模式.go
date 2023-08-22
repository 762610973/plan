package main

import "fmt"

// !代理覆盖自己原本的需求, 且可以额外做其他事情

type goods struct {
	Kind string
	Fact bool
}

var _ shopping = (*beijingShopping)(nil)
var _ shopping = (*shanghaiShopping)(nil)

type shopping interface {
	buy(g goods)
}

type beijingShopping struct{}

func (b beijingShopping) buy(g goods) {
	fmt.Println("beijing shopping", g.Kind)
}

type shanghaiShopping struct{}

func (s shanghaiShopping) buy(g goods) {
	fmt.Println("shanghai shopping", g.Kind)
}

// ! 代理
type proxy struct {
	shopping shopping
}

// 接受一个实现了shopping接口的具体对象
func newProxy(s shopping) shopping {
	return proxy{s}
}

func (p proxy) buy(g goods) {
	if p.distinguish(g) {
		p.shopping.buy(g)
	}
}

// 扩展能力, 辨别真伪
func (p proxy) distinguish(g goods) bool {
	return g.Fact
}

func main() {
	g := goods{
		Kind: "golang",
		Fact: true,
	}
	var s shopping
	s = new(beijingShopping)
	s.buy(g)
	var p shopping
	p = newProxy(s)
	p.buy(g)
}
