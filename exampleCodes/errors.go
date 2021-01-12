package main
import(
	"errors"
	"fmt"
	"os"
	"strconv"
)
func main(){
	
	if len(os.Args) == 1{
		fmt.Println("Please give one or more floats.")
		os.Exit(1)
	}
	arguments := os.Args
	fmt.Println("num: ",len(arguments))
	var err error = errors.New("An error")//1.에러를 정의해준다.
	k:=1
	var n float64
	
	for err!=nil{//에러가 발생하면 err!=nil이 true이므로 계속 돌아간다.
		fmt.Println("???: ",err!=nil)
		if k >= len(arguments){
			//k가 매개변수 개수를 초과하면 더 이상 float 타입의 매개변수가 존재하지 않음을 알 수 있음. 종료!
			fmt.Println("None of the arguments is a float!")
			return
		}
		
		n,err = strconv.ParseFloat(arguments[k],64)
		//ParseFloat으로 해당 매개변수가 float인지 아닌지를 판단한다.
		k++
	}
	
	//초기조건 : err가 정의되어있으므로 err!=nil 은 true이며 해당 루프는 작동한다.
	//유지조건 : ParseFloat은 err을 검출한다. 만약 float이 아니라면 err는 재정의되고
	//해당 루프는 유지된다.
	//종료조건 : 매개변수의 개수를 초과했을때 해당 루프는 끝이 난다. 또는 err가 없을때 종료된다.
	
	min, max := n,n
	
	for i:=2; i<len(arguments);i++{
		n,err = strconv.ParseFloat(arguments[i],64)
		if err == nil{
			if n<min{
				min=n
			}
			if n>max{
				max=n
			}
		}
	}
	fmt.Println(min)
	fmt.Println(max)
}