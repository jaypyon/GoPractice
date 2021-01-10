package main

import(
	"fmt"
	"github.com/mactsouk/go/simpleGitHub"
)
func main(){
	fmt.Println(simpleGitHub. AddTwo (5,6))
}

/*오류가 난 이유는 이 패키지가 머신에 설치돼 있지 않기 때문이다. 
먼저 커맨드를 실행해서 패키지를 다운로드한다.*/
//go get -v packagename or Url
//경로 : ~/go/src/github.com/mactsouk/go/simpleGitHub