package main

import "fmt"

/*
* 空接口any、interface{}的类型集是所有类型的集合
* 所以像int、string、strcut{}、MyStruct、func foobar()、chan int、map[int]string、[]int等等都在空接口的类型集合中
 */

/*
* 近似元素~T的类型集合是所有底层类型为T的所有类型的集合
* 一个非接口类型的类型集合就是只包含这个类型的类型集合, 比如int的类型集合只包含int这样一个元素
* 联合元素t1|t2|…|tn的类型集合是这些联合元素类型集合的并集
* 一个方法的类型集合是定义这个方法的所有类型,也就是只要某个类型的方法集包含这个方法
	* 那么它就属于这个方法的类型集合 比如接口中有String() string这样一个方法
	* 那么所有实现这个方法的类型都属于String() string定义的类型集合, 比如net.IP
*/

/*
- 只有包含类型元素的接口才定义了特定类型(可能是空的类型)
* 更准确地说，对于给定的接口I，特定类型的集合对应于该接口代表的类型集合𝑅，这里要求𝑅是非空且有限的。
* 否则，如果𝑅为空或无限，则接口没有特定类型
	* 联合类型的代表类型是无限的,没有特定类型
! 一个接口即使类型为空, 它的特定类型集合可能不为空
! 一个接口即使有有限的特定类型, 它的类型集合也可能是无限的
- 特定类型主要用于判断类型参数是否支持索引,还用做类型转化定义上
*/

func main() {
	//var m  myInt = 3
	m := myInt(3)
	m.name()
	testTwo(m)
	m2 := myInt2{3}
	testTwo(m2)
	range1([]byte{1, 2, 3}, 1, 2)
}

type two interface {
	~int | myInt2
	name() int
}

type myInt int

func (m myInt) name() int {
	return int(m)
}

func testTwo[T two](arg T) {
	arg.name()
}

type myInt2 struct {
	val int
}

func (m myInt2) name() int {
	return m.val
}

type twoPlus interface {
	two
	string | float64
}

func testTwoPlus[T two](arg T) {
	arg.name()
}

type myByte []byte

func range1[T []byte | myByte | string](x T, y, j int) {
	fmt.Println(len(x))
	fmt.Println(x[y:j])

}

func _[T interface{ []byte | myByte }](x T, y, j int) {
	fmt.Println(x[y:j])
}

type rangeType interface {
	[]byte | map[int]byte | string
}

// !报错,参考 下面的说明
/*
func rangeIt[R rangeType](r R) {
	for i, v := range r {
		fmt.Println(i, v)
	}
}
*/

/*
- 一个接口T要被成为结构化的(structural),需要满足下面的条件之一:
	* 存在一个单一的类型U,它是T的类型集合中的每一个元素相同的底层类型
	* T的类型集合只包含chan类型，并且它们的元素类型都是E
	* 所有的chan的方向包含相同的方向(不一定要求完全相同)
- 结构化类型包含一个结构类型，根据上面的条件不同，结构类型可能是:
	* 类型U, 或者
	* 如果T只包含双向chan的话，结构类型为chan E,否则可能是chan<- E或者<-chan E
*/

/*
- 如果a的类型是类型参数P的话
	* P必须有特定类型
	* 如果P的特定类型包含string类型，那么不能给a[x]赋值(字符串是不可变的)
	* P的所有特定类型必须相同
	* 如果P的特定类型包含map类型的话, 那么它的所有特定类型必须是map,而且所有的key的类型是相同的
- 特定类型还用作类型转化定义上, 对于一个变量x,如果它的类型是V, 要转换成的类型是T
	* V的每一个特定类型的值都可以转换成T的每一个特定类型
	* 只有V是类型参数，T不是，那么V的每一个特定类型的值都可以转换成T
	* 只有T是类型参数，x可以转换成T的每一个特定类型
*/

type t1 interface{}
type t2 interface{ int }
type t3 interface{ ~string }       // 特定类型是string
type t4 interface{ float64 | any } // 没有特定类型
type t5 interface {
	int
	string
} // 没有特定类型
type t6 interface {
	int
	any
} // 特定类型是int,int和any的交集
