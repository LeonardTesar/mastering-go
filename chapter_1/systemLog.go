package chapter_1

/** not included for windows bcs package is frozen - hence not further explored

import (
	"log"
	"log/syslog"
)


func systemLog() {
	sysLog, err := syslog.New(syslog.LOG_SYSLOG, "systemLog.go")
	if err != nil {
		log.Println(err)
		return
	} else {
		log.SetOutput(sysLog)
		log.Print("Everything is fine!")
	}
}
*/
