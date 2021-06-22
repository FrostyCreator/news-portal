package logger

import (
	"fmt"
	"github.com/FrostyCreator/news-portal/internal/utils"
	"io"
	"log"
	"os"
	"time"
)

const (
	logDir = "log/"
	lInfo  = "[INFO] "
	lWarn  = "[WARN] "
	lError = "[ERROR] "
	lFatal = "[FATAL] "
	lPanic = "[PANIC] "
)

var (
	infoLog  *log.Logger
	warnLog  *log.Logger
	errLog   *log.Logger
	fatalLog *log.Logger
	panicLog *log.Logger
)

func InitLogger() error {
	exists, err := utils.DirExists(logDir)
	if err != nil {
		return err
	}
	if !exists {
		if err = os.Mkdir(logDir, 0777); err != nil {
			return err
		}
	}

	fileName := fmt.Sprintf("%s%s_runlog.txt", logDir, time.Now().Format("2006-01-02_15:04:05"))
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	mw := io.MultiWriter(os.Stdout, file)

	infoLog = log.New(mw, lInfo, log.Ldate|log.Ltime)
	warnLog = log.New(mw, lWarn, log.Ldate|log.Ltime)
	errLog = log.New(mw, lError, log.Ldate|log.Ltime)
	fatalLog = log.New(mw, lFatal, log.Ldate|log.Ltime)
	panicLog = log.New(mw, lPanic, log.Ldate|log.Ltime)

	return nil
}

func LogInfo(v ...interface{}) {
	infoLog.Print(v)
}

func LogInfof(format string, v ...interface{}) {
	infoLog.Printf(format, v)
}

func LogWarn(v ...interface{}) {
	warnLog.Print(v)
}

func LogWarnf(format string, v ...interface{}) {
	warnLog.Printf(format, v)
}

func LogError(v ...interface{}) {
	errLog.Print(v)
}

func LogErrorf(format string, v ...interface{}) {
	errLog.Printf(format, v)
}

func LogFatal(v ...interface{}) {
	fatalLog.Print(v)
	os.Exit(1)
}

func LogFatalf(format string, v ...interface{}) {
	fatalLog.Printf(format, v)
	os.Exit(1)
}

func LogPanic(v ...interface{}) {
	panicLog.Print(v)
	panic(fmt.Sprint(v))
}

func LogPanicf(format string, v ...interface{}) {
	panicLog.Printf(format, v)
	panic(fmt.Sprintf(format, v))
}
