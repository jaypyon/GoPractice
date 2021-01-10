package main

import (
	"fmt"
)

func main()
{
	fmt.Println("Go has strict rules for curly braces!")
}

//중괄호 curly braces 사용 역시 굉장히 엄격하다.
/*
컴파일러는 코드에서 필요하다고 판단되는 지점에 세미콜론을 집어넣는다.
따라서 여는 괄호 로 문장을 시작하면 60 컴파일러는 그 이전 문장의 끝(func main)에
세미콜론을 넣기 때문에 위와 같은 에러 메시지가 발생하는 것이다.
*/