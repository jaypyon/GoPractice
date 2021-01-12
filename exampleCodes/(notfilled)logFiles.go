package main
import(
	"log"
	"fmt"
	"log/syslog"
	"os"
	"path/filepath"
)

func main(){
	progamName := filepath.Base(os.Args[0])
	sysLog,err := syslog.New(syslog.Log_NOTICE|syslog.LOG_MAIL,programName)
	//로그 종류는 메일이고 수준은 노티스인 메시지
		/*class LogLevel
		{
			const EMERGENCY = 'emergency';
			const ALERT     = 'alert';
			const CRITICAL  = 'critical';
			const ERROR     = 'error';
			const WARNING   = 'warning';
			const NOTICE    = 'notice';
			const INFO      = 'info';
			const DEBUG     = 'debug';
		}*/
	if err != nil{
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
	}
	log.Println("LOG_NOTICE + LOG_MAIL")
	sysLog, err = syslog.New(syslog.LOG_MAIL,"Some program!")
}