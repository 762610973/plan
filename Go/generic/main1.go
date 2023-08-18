package main

import (
	"fmt"
	"io"
)

func main() {
	fmt.Println("💘 study Golang's generic 💘")
	fmt.Println("🎺 ~T的类型T必须是底层类型自己, 而且不能是接口类型")
	fmt.Println("🎺 联合(union)类型元素不能是类型参数")
	fmt.Println("🎺 联合(union)类型元素的非接口元素必须是两两不相交")
	fmt.Println("🎺 非传统接口只能用作类型约束, 或者其它约束接口的元素")
}

// 类型集,类型元素包含类型T,用作类型约束
type generic1 interface {
	int
}

// 类型集,类型元素包含近似类型T
type generic2 interface {
	~int
}

// 类型集,类型元素包含联合类型(A|B|C~D)
/*
* 联合类型的元素不能是类型参数
! interface{ K }中K是类型参数
func I1[K any, V interface{ K }]() {}

! 错误, interface{ nt | K }中K是类型参数
func I2[K any, V interface{ int | K }]() {}
*/

type generic3 interface {
	int
	string
}

// 好像没啥用, 目前不知道有啥用
type generic4 interface {
	string
	study()
}

type generic5 interface {
	any
}

var _ = generic5(3)
var _ = generic5("generic")

// - 联合(union)类型元素的非接口元素必须是两两不相交

//! 错误！int和~int相交
//* Goland不会提示,但是编译时会报错
//func I4[K any, V interface{ int | ~int }]() {}

// MyInt
// 下面的定义没有问题。因为int和MyInt是两个类型，不相交
type MyInt int

// I5 不相交
func I5[K any, V interface{ int | MyInt }]() {}

// ! I6 错误! int和~MyInt相交, 交集是int
// Goland会直接提示错误
//func I6[K any, V interface{ int | ~MyInt }]() {}

type MyInt2 = int

// ! 编译报错, 提示重叠
// Goland不会提示
//func I7[K any, V interface{ int | MyInt2 }]() {}

var (
	// 以下编译没问题
	_ interface{}
	_ interface{ m() }
	_ interface{ io.Reader }
	_ interface {
		io.Reader
		io.Writer
	}
	// 以下不能编译, 接口不能用作变量实例类型
	//_ interface{ int }
	//_ interface{ ~int }
	//_ interface{ MyInt }
	/*A interface {
		int
		m()
	}*/
	// 可以编译
	_ struct{ i int }
	// 下面一行不能编译,因为~int不能作为字段的类型
	//_ struct{ i ~int }
	// 下面一行不能编译，因为constraints.Ordered只能用作类型约束
	//_ struct{ i constraints.Ordered }
	// 行能够编译，是接口类型，并且类型元素也是普通接口
	_ interface{ any }
	// 不能编译
	/*_ interface {
		interface {
			any
			m()
		}
	}*/
	// 不能编译, 因为接口不属于普通接口, 而是类型约束用作类型集
	/*
		_ interface {
			interface {
				int | ~int
				m()
			}
		}
	*/
)
