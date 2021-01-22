package main

import (
	"fmt"
)

type aStructure struct {
	name   string
	height int
	weight int
	inbody float64
}

func createStructureA(n string, h, w int, i float64) *aStructure {
	return &aStructure{n, h, w, i}
}
func main() {
	p := fmt.Println
	s1 := createStructureA("Chiwon", 175, 67, 21.3)
	p(s1) //s1은 현재 포인터임을 알 수 있다.
	p(s1.name)
	p(*s1)

	b := make([]aStructure, 4)
	p(b)

	c := new([]aStructure)
	p(c)
}
