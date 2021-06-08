package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
)

/*
 * 日志输出管理：按照日志等级进行输出，等级越高，输出的日志越少
 * 日志时间：精确到毫秒
 * 输出格式：yyyy/mm/dd hh:MM:ss.microsecond xxx/../xx.go:86: [DEBUG] your log content.
 */

var (
	marsLog      = log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds)
	marsLogLevel = DEBUG
	projectName  = "mars"
	maxHeader    = 128
)

type LogLevel int

// 日志等级定义
const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

var LogLevelStr = []string{
	DEBUG: "[DEBUG]",
	INFO:  "[INFO]",
	WARN:  "[WARN]",
	ERROR: "[ERROR]",
}

const calldepth = 3

// 设置日志等级，默认为 DEBUG
func SetLogLevel(level LogLevel) {
	marsLogLevel = level
}

func Debug(v ...interface{}) {
	output(DEBUG, v...)
}

func Debugf(format string, v ...interface{}) {
	outputf(DEBUG, format, v...)
}

func Info(v ...interface{}) {
	output(INFO, v...)
}

func Infof(format string, v ...interface{}) {
	outputf(INFO, format, v...)
}

func Warn(v ...interface{}) {
	output(WARN, v...)
}

func Warnf(format string, v ...interface{}) {
	outputf(WARN, format, v...)
}

func Error(v ...interface{}) {
	var buf [2 << 10]byte
	v = append(v, string(buf[:runtime.Stack(buf[:], false)]))
	output(ERROR, v...)
}

func Errorf(format string, v ...interface{}) {
	var buf [2 << 10]byte
	v = append(v, string(buf[:runtime.Stack(buf[:], false)]))
	outputf(ERROR, format, v...)
}

// 致命错误日志，会调用os.Exit(1)退出程序
func Fatal(v ...interface{}) {
	marsLog.Fatal(v...)
}

// 致命错误日志，会调用os.Exit(1)退出程序
func Fatalf(format string, v ...interface{}) {
	marsLog.Fatalf(format, v...)
}

func checkLevel(level LogLevel) bool {
	return level >= marsLogLevel
}

func formatHeader(level LogLevel, data string) string {
	// 格式化代码文件路径
	_, file, line, ok := runtime.Caller(calldepth)
	if !ok {
		file = "???"
		line = 0
	}

	i := strings.Index(file, projectName)
	if i != -1 {
		file = file[i:]
	}

	header := strings.Builder{}
	header.Grow(maxHeader)

	// 过早获取时间戳，会导致与log写入间隔拉长，可能会出现log不按时间戳顺序显示（毫秒级）
	//header.WriteString(fmt.Sprintf("%s.%06d", time.Now().Format("2006-01-02 15:04:05"), time.Now().Nanosecond()/1e3))

	header.WriteString(" /")
	header.WriteString(file)
	header.WriteString(":")
	header.WriteString(strconv.Itoa(line))
	header.WriteString(" ")
	header.WriteString(LogLevelStr[level])
	header.WriteString(" ")
	header.WriteString(data)
	return header.String()
}

func output(level LogLevel, v ...interface{}) {
	if checkLevel(level) {
		marsLog.Output(calldepth, formatHeader(level, fmt.Sprintln(v...)))
	}
}

func outputf(level LogLevel, format string, v ...interface{}) {
	if checkLevel(level) {
		marsLog.Output(calldepth, formatHeader(level, fmt.Sprintf(format, v...)))
	}
}
