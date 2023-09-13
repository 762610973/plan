package main

import (
	"fmt"
	"log/slog"
	"maps"
	"strings"
)

/*
* 创建型模式的一种, 特点在于通过"复制"一个已经存在的实例来返回新的实例, 而不是新建实例.
* 被复制的实例就是我们所称的"原型"
 */

type val struct {
	v int
}

func main() {
	var s = "abc"
	// 当值保留一个大字符串中的一小部分子串时, 非常重要
	clone := strings.Clone(s[:1])
	slog.Info(clone)
	v := &val{v: 1}
	m1 := map[int]*val{1: v}
	// maps.Clone是一个浅克隆
	m2 := maps.Clone(m1)
	fmt.Println(m1[1])
	fmt.Println(m2[1])
	v.v = 100
	fmt.Println(m1[1])
	fmt.Println(m2[1])
}

// https://github.com/jinzhu/copier
// https://github.com/switchupcb/copygen
// https://github.com/jmattheis/goverter
