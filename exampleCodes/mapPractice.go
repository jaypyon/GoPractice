package main

import (
	"fmt"
)

func main() {
	newMap := map[string]int{
		"key":   530,
		"index": 321,
	}

	fmt.Println(newMap)
	delete(newMap, "key")
	fmt.Println(newMap)
}
