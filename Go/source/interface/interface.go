package main

// 声明式校验
var _ Dayer = &Today{}

// 不能这样使用, 指针实现接口
// var _ Dayer = Today{}
// 类型转换形式
var _ Dayer = (*Today)(nil)

// 声明式, 由于Tomorrow没有使用指针实现接口, 所以可以直接使用
var _ Dayer = Tomorrow{}

var _ Dayer = &Tomorrow{}

// 值接收器实现的接口也可以这样使用
var _ Dayer = (*Tomorrow)(nil)

func main() {

}

type Today struct{}

func (*Today) Date() {

}

type Tomorrow struct{}

func (Tomorrow) Date() {

}

type Dayer interface {
	Date()
}
