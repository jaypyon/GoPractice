package main
import(
	"io"
	"bufio"//파일 입출력에 관한 것
	_"fmt"
	"os"//os.Args(커맨드라인 인수)에 접근하기 위해 자주 쓰인다.
)

func main(){
	var f *os.File
	f= os.Stdin
	defer f.Close()
	
	scanner :=bufio.NewScanner(f)
	for scanner.Scan(){
		io.WriteString(os.Stdout,">"+scanner.Text()+"\n")
		
	}
}