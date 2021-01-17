package main

import "fmt"

func main() {
	var arrayThree = [2][2][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}
	fmt.Println(arrayThree[1])
	fmt.Println(arrayThree[1][1])
	fmt.Println(arrayThree[1][1][0])
	fmt.Println(arrayThree[1][1][1])
}
