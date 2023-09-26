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
	fmt.Println(tf.Name())
	// 类型方法数量
	fmt.Println(tf.NumMethod())
	// 返回这个Type的特定类型
	fmt.Println(tf.Kind())
	fmt.Println(tf.Kind().String())
	// 返回这个类型的大小
	fmt.Println(tf.Size())
	// 返回此类型的内存对齐大小
	fmt.Println(tf.Align())
	// 返回此类型的字符串表达形式
	fmt.Println(tf.String())
	// 没有方法不能调用Method()
	//fmt.Println(tf.Method(0))
	// 非结构体类型不能调用NumField
	//fmt.Println(tf.NumField())
	// 非map类型, 不能调用Key()
	//fmt.Println(tf.Key())
}

func TestMap(t *testing.T) {
	m := map[int]int{}
	m[1] = 1
	tf := reflect.TypeOf(m)
	// 返回key的类型
	fmt.Println(tf.Key())
	fmt.Println(tf.Key().Kind() == reflect.Int)
	fmt.Println(tf.Kind())
}
