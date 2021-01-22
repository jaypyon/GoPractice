package main

import (
	"fmt"
)

type Digit int

const (
	ONE Digit = iota
	TWO
	THREE
)
const (
	p1 Digit = 1 << iota
	p2
	p3
	p4
)

func main() {
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
}
