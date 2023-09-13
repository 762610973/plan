package main

import "fmt"

// beverage 抽象类,
type beverage interface {
	boilWater()
	brew()
	pourInCup()
	addTings()
}

type template struct {
	b beverage
}

func (t *template) makeBeverage() {
	if t == nil {
		return
	}
	// 固定流程,具体是执行子类具有的
	t.b.boilWater()
	t.b.brew()
	t.b.pourInCup()
	t.b.addTings()
}

type makeCoffee struct {
	template
}

func (m makeCoffee) boilWater() {
	fmt.Println("boil water")

}

func (m makeCoffee) brew() {
	fmt.Println("brew")
}

func (m makeCoffee) pourInCup() {
	fmt.Println("pour in cup")
}

func (m makeCoffee) addTings() {
	fmt.Println("add tings")
}

func newMakeCoffee() *makeCoffee {
	mc := new(makeCoffee)
	// 给接口赋值
	mc.b = mc
	return mc
}

func main() {
	mc := newMakeCoffee()
	mc.makeBeverage()
}
