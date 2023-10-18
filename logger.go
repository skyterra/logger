package logger

import (
	"context"
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
	ins         = log.New(os.Stdout, "", log.Lmsgprefix)
	logLevel    = debug
	srcFolder   = ""
	projectName = "♫"
)

const RequestID = "sk-request-id"

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

func Init(level string, srcFolder string, projectName string) {
	SetLevel(level)
	SetSrcFolder(srcFolder)
	SetProjectName(projectName)
}

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

func Debug(ctx context.Context, v ...interface{}) {
	output(ctx, debug, v...)
}

func Debugf(ctx context.Context, format string, v ...interface{}) {
	outputf(ctx, debug, format, v...)
}

func Info(ctx context.Context, v ...interface{}) {
	output(ctx, info, v...)
}

func Infof(ctx context.Context, format string, v ...interface{}) {
	outputf(ctx, info, format, v...)
}

func Warn(ctx context.Context, v ...interface{}) {
	output(ctx, warn, v...)
}

func Warnf(ctx context.Context, format string, v ...interface{}) {
	outputf(ctx, warn, format, v...)
}

func Error(ctx context.Context, v ...interface{}) {
	var buf [2 << 10]byte
	v = append(v, string(buf[:runtime.Stack(buf[:], false)]))
	output(ctx, error, v...)
}

func Errorf(ctx context.Context, format string, v ...interface{}) {
	var buf [2 << 10]byte
	v = append(v, string(buf[:runtime.Stack(buf[:], false)]))
	outputf(ctx, error, format, v...)
}

// 致命错误日志，会调用os.Exit(1)退出程序
func Fatal(v ...interface{}) {
	ins.Fatal(v...)
}

// 致命错误日志，会调用os.Exit(1)退出程序
func Fatalf(format string, v ...interface{}) {
	ins.Fatalf(format, v...)
}

func checkLevel(level int) bool {
	return level >= logLevel
}

func formatHeader(ctx context.Context, level int, data string) string {
	requestID, _ := ctx.Value(RequestID).(string)
	if requestID != "" {
		requestID = fmt.Sprintf(" [%s]", requestID)
	}

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
	header.WriteString(fmt.Sprintf(" %s %s:%s [%s]%s$ ", projectName, file, strconv.Itoa(line), levelName[level], requestID))
	header.WriteString(data)
	return header.String()
}

func output(ctx context.Context, level int, v ...interface{}) {
	if checkLevel(level) {
		ins.Output(calldepth, formatHeader(ctx, level, fmt.Sprintln(v...)))
	}
}

func outputf(ctx context.Context, level int, format string, v ...interface{}) {
	if checkLevel(level) {
		ins.Output(calldepth, formatHeader(ctx, level, fmt.Sprintf(format, v...)))
	}
}
