package main
import(
	"os"
	"io"
)
func main(){
	myString:=""
	arguments:= os.Args
	
	if len(arguments)==1{
		myString="Please give me one argument!"
	}else {
		myString = arguments[1]
	}
	
	io.WriteString(os.Stdout,"Hello, World!")
	io.WriteString(os.Stderr,myString)
	io.WriteString(os.Stderr,"\n")
}


/*표준 출력과 표준 에러의 결과를 모두 같은파일에 저장하고 싶다면, 표준 에러에 대한 파일 디스크립터(2)를 표준 출력에 대한 파일
디스크립터(1)로 리디렉션하면 된다. 예를 들면 다음과 같다. 유닉스 시스템에서는 이러한 표현을 꽤 자주 사용한다.
(명령어뒤에) >/tmp/output 2>&1
*/