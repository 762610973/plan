package main

import "fmt"

// !适配器模式: 将一个类的接口转换成客户希望的另一个接口。适配器模式使得原本由于接口不兼容而不能一起工作的那些类可以

// 适配的目标
type v5 interface {
	use5v()
}

// 被适配的角色,适配者
type v220 struct{}

func (v v220) use220v() {
	fmt.Println("使用220v的电压")
}

// 适配器类
type v220Adapter struct {
	v220 v220
}

// 构造适配器, 接受一个被适配的角色,在这里220v是被适配的
func newV220Adapter(v220 v220) v220Adapter {
	return v220Adapter{
		v220: v220,
	}
}

func (v v220Adapter) use5v() {
	v.v220.use220v()
}

// 业务类
type phones struct {
	v v5
}

// 接受一个实现了v5接口的适配器类
func newPhone(v v5) phones {
	return phones{
		v: v,
	}
}

func (p phones) charge() {
	fmt.Println("phone进行了充电")
	p.v.use5v()
}

func main() {
	// 被适配的角色
	a := v220{}
	p := newPhone(newV220Adapter(a))
	p.charge()
}
