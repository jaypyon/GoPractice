package main

import (
	"fmt"
)

func retThree(x int) (int, int, int) {
	return x, x * x, x * x * x
}

func main() {
	fmt.Println(retThree(3))
}
