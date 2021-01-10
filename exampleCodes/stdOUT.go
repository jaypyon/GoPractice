package main

import(
	"io"
	"os"
)
func main(){
	myString:=""
	arguments:= os.Args
	if len(arguments)==1{
		myString = "Please give me one argument!"
	} else{
		myString =arguments[1]
	}
	io.WriteString(os.Stdout,myString)
	io.WriteString(os.Stdout,"\n")
}
//io.WriteString 함수의 첫번째 매개변수는 io.write타입이다.
//두번째 매개변수는 byte slice로 지정해야하지만 string으로 사용해도 무관하다.

//io.WriteString 함수로 os.Stdout에 myString을 넣어준다.