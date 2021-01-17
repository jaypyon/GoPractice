package main

import "fmt"

func main() {
	var i int = 0
	anExpression := true
	for ok := true; ok; ok = anExpression {
		if i > 10 {
			anExpression = false
			//11에서 false로 바뀌면
		}
		fmt.Println(i, " ")
		i++
	}
	fmt.Println()

	var value = [5]int{0, 1, -1, 2, -2}
	for i, val := range value {
		fmt.Println("index: ", i, "value:", val)
	}
}
