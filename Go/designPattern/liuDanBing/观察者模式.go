package main

import (
	"fmt"
)

// 抽象的观察者
type listener interface {
	onTeacherComing() // 观察者得到通知后要触发的具体工作
}

type notifier interface {
	addListener(l listener)
	removeListener(l listener)
	notify()
}

// 观察者具体的学生

type stu struct {
	badTing string
}

func (s stu) doBadThing() {
	fmt.Println("do bad thing")
}

func (s stu) onTeacherComing() {
	fmt.Println("stu stop", s.badTing)
}

// 通知者

type monitor struct {
	listenerList []listener
}

func (m monitor) addListener(l listener) {
	m.listenerList = append(m.listenerList, l)
}

func (m monitor) removeListener(l listener) {
	for index, listen := range m.listenerList {
		if listen == l {
			m.listenerList = append(m.listenerList[:index], m.listenerList[index+1:]...)
			break
		}
	}
}

func (m monitor) notify() {
	for _, lis := range m.listenerList {
		lis.onTeacherComing()
	}
}

func main() {
	s := stu{badTing: "抄作业"}
	var m notifier
	m = new(monitor)
	m.addListener(s)
	s.doBadThing()
	m.notify()
}
