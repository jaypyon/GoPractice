package main

import "fmt"

func main() {
	aSlice := []int{1, 2, 3, 4, 5}
	bSlice := make([]int, 4)
	fmt.Println(aSlice)
	fmt.Println(bSlice)
	fmt.Println(cap(aSlice), len(aSlice))
	fmt.Println(cap(bSlice), len(bSlice))
	copy(bSlice, aSlice)

	fmt.Println(aSlice)
	fmt.Println(bSlice)
	fmt.Println(cap(aSlice), len(aSlice))
	fmt.Println(cap(bSlice), len(bSlice))

	cSlice := make([][]int, 3)
	cSlice = nil
	fmt.Println(cSlice)
}
