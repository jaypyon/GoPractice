package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Epoch time :", time.Now().Unix())
	t := time.Now()
	fmt.Println(t, " %hello% ", t.Format(time.RFC3339))
	fmt.Println(t.Weekday(), t.Day(), t.Month(), t.Year())

	time.Sleep(time.Second)
	t1 := time.Now()
	fmt.Println("Time difference:")

	fmt.Println(" %hello% :", t1)
	formaT := t.Format("01 January 2000")
	fmt.Println(formaT)
}

//2021-01-18T17:10:12+09:00 //RFC3999
//2021-01-18 17:10:12.8117819 +0900 KST m=+0.008024901 //time.Now()
//Monday 18 January 2021
//t.Weekday(), t.Day(), t.Month(), t.Year()
