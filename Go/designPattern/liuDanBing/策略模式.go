package main

import "fmt"

type weaponStrategy interface {
	useWeapon()
}

// 具体策略

type ak47 struct{}

func (a ak47) useWeapon() {
	fmt.Println("ak47")
}

type knife struct{}

func (k knife) useWeapon() {
	fmt.Println("knife")
}

type hero struct {
	strategy weaponStrategy
}

func (h *hero) setWeaponStrategy(w weaponStrategy) {
	h.strategy = w
}

func (h *hero) fight() {
	h.strategy.useWeapon()
}

func main() {
	h := new(hero)
	h.setWeaponStrategy(ak47{})
	h.fight()

	h.setWeaponStrategy(knife{})
	h.fight()
}
