package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

/*
 * 日志输出管理：按照日志等级进行输出，等级越高，输出的日志越少
 * 日志时间：精确到毫秒
 * 输出格式：yyyy/mm/dd hh:MM:ss.xxx xxx/../xx.go:86: [debug] your log content.
 */

var (
	lanhuLog    = log.New(os.Stdout, "", log.Lmsgprefix)
	logLevel    = debug
	srcFolder   = ""
	projectName = "UnknownProject"
)

// 日志等级定义
const (
	debug = iota
	info
	warn
	error
)

var levelName = []string{
	debug: "DEBUG",
	info:  "INFO",
	warn:  "WARN",
	error: "ERROR",
}

const (
	calldepth = 3   // 错误日志堆栈深度
	maxHeader = 128 // 日志最大长度
)

// 设置日志等级 "debug", "info", "warn", "error"
func SetLevel(level string) {
	level = strings.ToUpper(level)

	switch level {
	case levelName[debug]:
		logLevel = debug
	case levelName[info]:
		logLevel = info
	case levelName[warn]:
		logLevel = warn
	case levelName[error]:
		logLevel = error
	default:
		logLevel = debug
	}
}

func SetProjectName(name string) {
	projectName = name
}

func SetSrcFolder(folderName string) {
	srcFolder = folderName
}

func Debug(v ...interface{}) {
	output(debug, v...)
}

func Debugf(format string, v ...interface{}) {
	outputf(debug, format, v...)
}

func Info(v ...interface{}) {
	output(info, v...)
}

func Infof(format string, v ...interface{}) {
	outputf(info, format, v...)
}

func Warn(v ...interface{}) {
	output(warn, v...)
}

func Warnf(format string, v ...interface{}) {
	outputf(warn, format, v...)
}

func Error(v ...interface{}) {
	var buf [2 << 10]byte
	v = append(v, string(buf[:runtime.Stack(buf[:], false)]))
	output(error, v...)
}

func Errorf(format string, v ...interface{}) {
	var buf [2 << 10]byte
	v = append(v, string(buf[:runtime.Stack(buf[:], false)]))
	outputf(error, format, v...)
}

// 致命错误日志，会调用os.Exit(1)退出程序
func Fatal(v ...interface{}) {
	lanhuLog.Fatal(v...)
}

// 致命错误日志，会调用os.Exit(1)退出程序
func Fatalf(format string, v ...interface{}) {
	lanhuLog.Fatalf(format, v...)
}

func checkLevel(level int) bool {
	return level >= logLevel
}

func formatHeader(level int, data string) string {
	// 格式化代码文件路径
	_, file, line, ok := runtime.Caller(calldepth)
	if !ok {
		file = "???"
		line = 0
	}

	i := strings.Index(file, srcFolder)
	if i != -1 {
		file = file[i:]
	}

	header := strings.Builder{}
	header.Grow(maxHeader)

	// 过早获取时间戳，会导致与log写入间隔拉长，可能会出现log不按时间戳顺序显示（毫秒级）
	header.WriteString(fmt.Sprintf("%s.%03d", time.Now().Format("2006/01/02-15:04:05"), time.Now().Nanosecond()/1e6))
	header.WriteString(fmt.Sprintf(" %s %s:%s [%s]$ ", projectName, file, strconv.Itoa(line), levelName[level]))
	header.WriteString(data)
	return header.String()
}

func output(level int, v ...interface{}) {
	if checkLevel(level) {
		lanhuLog.Output(calldepth, formatHeader(level, fmt.Sprintln(v...)))
	}
}

func outputf(level int, format string, v ...interface{}) {
	if checkLevel(level) {
		lanhuLog.Output(calldepth, formatHeader(level, fmt.Sprintf(format, v...)))
	}
}
