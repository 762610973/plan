package main

import (
	"fmt"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"time"
)

var m2 = make([]*myBirthday, 0, 1024)

func main() {
	go birthday()
	go randMap()
	_ = http.ListenAndServe(":6060", nil)

}
func birthday() {
	for {
		fmt.Println("祝我自己生日快乐~~")
		time.Sleep(time.Second)
	}
}

func randMap() {
	m := make(map[int64]int64)
	for {
		m[rand.Int63()] = rand.Int63n(13)
		b := &myBirthday{
			name: "xiaoLiang",
			age:  23,
		}
		m2 = append(m2, b)
		b.congratulate()
		time.Sleep(time.Millisecond * 100)
		fmt.Println("the len of map", len(m))
	}
}

type myBirthday struct {
	name string
	age  int8
}

func (m myBirthday) congratulate() {
	fmt.Printf("祝%s%d岁生日快乐\n", m.name, m.age)
}
