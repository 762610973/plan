package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInt(t *testing.T) {
	var a int
	tf := reflect.TypeOf(a)
	// 类型名
	fmt.Println("name: ", tf.Name())
	// 类型方法数量
	fmt.Println(tf.NumMethod())
	// 返回这个Type的特定类型
	fmt.Println(tf.Kind())
	fmt.Println(tf.Kind().String())
	// 返回这个类型的大小
	fmt.Println("size: ", tf.Size())
	// 返回此类型的内存对齐大小
	fmt.Println("align: ", tf.Align())
	fmt.Println("bits: ", tf.Bits())

	// 返回此类型的字符串表达形式
	fmt.Println(tf.String())
	// 没有方法不能调用Method()
	//fmt.Println(tf.Method(0))
	// 非结构体类型不能调用NumField
	//fmt.Println(tf.NumField())
	// 非map类型, 不能调用Key()
	//fmt.Println(tf.Key())
	// 是否可比较
	fmt.Println("comparable: ", tf.Comparable())
	fmt.Println(tf.PkgPath())
}

func TestMap(t *testing.T) {
	m := make(map[int]string)
	m[1] = "one"
	tf := reflect.TypeOf(m)
	// 返回key的类型
	fmt.Println("type of string: ", tf.String())
	fmt.Println("map's key string: ", tf.Key().String())
	fmt.Println("map's key name: ", tf.Key().Name())
	fmt.Println("map's key kind: ", tf.Key().Kind())
	fmt.Println(tf.Key().Kind() == reflect.Int)
	fmt.Println("kind: ", tf.Kind())
	fmt.Println("kind string: ", tf.Kind().String())
	fmt.Println("name: ", tf.Name())
	fmt.Println("num method: ", tf.NumMethod())
	fmt.Println("map's string: ", tf.String())
	fmt.Println("align: ", tf.Align())
	fmt.Println("size: ", tf.Size())
	fmt.Println("comparable: ", tf.Comparable())
	// 返回val的类型
	fmt.Println(tf.Elem())
	fmt.Println(tf.PkgPath())
}

func TestSlice(t *testing.T) {
	s := []int{1}
	tf := reflect.TypeOf(s)
	fmt.Println(tf.String())
	fmt.Println("name: ", tf.Name())
	fmt.Println("kind: ", tf.Kind())
	fmt.Println("kind string: ", tf.Kind().String())
	fmt.Println("kind string: ", tf.PkgPath())
}
