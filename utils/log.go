package utils

import (
	"log"
	"os"
)

var consoleLog *log.Logger
var fileLog *log.Logger

var enableConsoleLog = false

func InitLog(enableConsoleLogParam bool) {
	enableConsoleLog = enableConsoleLogParam

	if enableConsoleLog {
		consoleLog = log.New(os.Stdout, "[Mini-JVM] ", log.Ldate|log.Ltime)
	}

}

func LogInfoPrint(v ...interface{}) {
	if enableConsoleLog {
		consoleLog.Print(v...)
	}
}

func LogInfoPrintln(v ...interface{}) {
	if enableConsoleLog {
		consoleLog.Println(v...)
	}
}

func LogInfoPrintf(format string, v ...interface{}) {
	if enableConsoleLog {
		consoleLog.Printf(format, v...)
	}
}

func LogErrorPrint(v ...interface{}) {
	createVmErrorLog()
	fileLog.Print(v...)

	LogInfoPrint(v...)
}

func LogErrorPrintln(v ...interface{}) {
	createVmErrorLog()
	fileLog.Println(v...)

	LogInfoPrintln(v...)
}

func LogErrorPrintf(format string, v ...interface{}) {
	createVmErrorLog()
	fileLog.Printf(format, v...)

	LogInfoPrintf(format, v...)
}

func createVmErrorLog() {
	if nil != fileLog {
		return
	}

	logFile, err := os.OpenFile("vm-error.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0766)
	if nil != err {
		panic(err)
	}

	fileLog = log.New(logFile, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
}
