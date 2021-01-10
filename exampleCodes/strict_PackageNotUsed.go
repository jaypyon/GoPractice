package main

import(
	"fmt"
	_"os"
)
func main(){
	fmt.Println("Hello there!")
}


//사용되지 않는 os라는 패키지가 import 되어있기 때문에 실행 되지 않는다.
//Go 언어는 이렇듯 패키지 사용에 대해 엄격한 규칙을 적용하고있다.
//사용되지 않는 패키지 앞에 _언더바를 붙이면 컴파일 에러 메시지가 발생하지 않는다.