package main

import "fmt"

type game struct {
	On    bool
	Ammo  int
	Power int
}

func (g *game) Shoot() bool {
	if g.On == false {
		return false
	} else if g.Ammo > 0 {
		g.Ammo -= 1
		return true
	} else {
		return false
	}
}

func (g *game) RideBike() bool {
	if g.On == false {
		return false
	} else if g.Power > 0 {
		g.Power -= 1
		return true
	} else {
		return false
	}
}
func main() {
	newOne := game{true, 1, 2}
	testStruct := &newOne
	fmt.Println(testStruct.RideBike())
}
