package main
import(
	"fmt"
	"strconv"//스트링을 산술 데이터 타입으로 변환
	"os"
)
func main(){
	if len(os.Args)==1{
		fmt.Println("Please give one or more floats")
		os.Exit(1)
	}
	
	arguments:=os.Args
	min,_:=strconv.ParseFloat(arguments[1],64)
	max,_:=strconv.ParseFloat(arguments[1],64)
	//n,_:=strconv.ParseFloat(arguments[1],64) 
	// 리턴한 값 중 첫 번째 값만 가져가고 두번째 값 무시한다. 에러값을 언더스코어(빈 식별자)에 할당한다.
	//에러 무시하는 것은 위험한 습관이다.
	
	for i:=2;i<len(arguments); i++{
		n,_:= strconv.ParseFloat(arguments[i],64)
		if n<min{
			min=n
		}
		if n>max{
			max=n
		}
	}
	fmt.Println("Min : ",min)
	fmt.Println("Max : ",max)
}