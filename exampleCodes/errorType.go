package main
import(
	"fmt"
	"os"
	"errors"
)
func returnError(a,b int) error{ // 함수의 선언, 매개변수 뒤에 리턴타입
	//현재 리턴타입은 error 
	
	if a == b {
		err:= errors.New("Error in returnError() function!")
		return err
	} else{
		return nil
	}
}
func main(){
	err:= returnError(1,2)
	if err == nil{
		fmt.Println("returnError() ended normally.")
	} else {
		fmt.Println(err)
	}
	
	err = returnError(10,10)
	if err == nil{
		fmt.Println("returnError() ended normally")
	} else {
		fmt.Println(err)
	}
	
	if err.Error() == "Error in returnError() function!"{
		fmt.Println("!!")
		os.Exit(10)
		/*Error() 함수 사용으로 errors 객체에 정의된 
		string 타입의 변수를 사용할 수 있다. */
	}
}