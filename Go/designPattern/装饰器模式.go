package main

import "fmt"

var (
	_ phone = (*huawei)(nil)
	_ phone = (*xiaomi)(nil)
)

func main() {
	var h phone
	h = huawei{}
	h.show()
	// 装饰
	var a phone
	a = newAdapter(h)
	a.show()
}

// !抽象层

type phone interface {
	show()
}

type adapter struct {
	phone phone
}

func newAdapter(p phone) phone {
	return adapter{phone: p}
}

func (a adapter) show() {
	// 调用被装饰的构件的原方法
	a.phone.show()
	// 装饰器额外的功能
	fmt.Println("this is adapter, add extra feature")
}

type huawei struct{}

func (h huawei) show() {
	fmt.Println("this is huawei")
}

type xiaomi struct{}

func (x xiaomi) show() {
	fmt.Println("this is xiaomi")
}
