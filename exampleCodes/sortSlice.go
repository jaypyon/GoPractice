package main

import (
	"fmt"
	"sort"
)

type aStructure struct {
	person string
	height int
	weight int
}

func main() {
	mySlice := make([]aStructure, 0)
	mySlice = append(mySlice, aStructure{"sukhan", 530, 80})
	mySlice = append(mySlice, aStructure{"yongjae", 185, 190})
	mySlice = append(mySlice, aStructure{"joonsung", 180, 90})

	fmt.Println("0: ", mySlice)

	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height > mySlice[j].height
	})

	fmt.Println(mySlice)
	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].weight > mySlice[j].weight
	})
	fmt.Println(mySlice)
	mySlice = append(mySlice, aStructure{"jaeinmoon", 123, 32})
	fmt.Println(mySlice)

}
