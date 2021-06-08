# st-logger

## 项目介绍
按照日志等级进行格式化输出

## 安装
```shell
go get github.com/skyterra/st-logger
```

## Example

```go
package main

import "github.com/skyterra/st-logger/logger"

func main() {
	logger.SetLogLevel(logger.DEBUG)
	logger.SetProjectName("your-project")
	logger.Debug("hello world")
}
```

```shell
2021/06/08 17:23:40.288792 /your-project/main.go:267 [DEBUG] hello world
```

## 日志等级
|ID |TAG  |描述|
|---|---  |---|
|0  |DEBUG|输出调试日志，通常用于打印调试信息
|1  |INFO |输出行为日志，通常用于追踪程序执行流程
|2  |WARN |输出警告日志，通常用于常规异常行为，如空指针判断
|3  |ERROR|输出错误日志，通常用于严重的异常，即，可能导致程序crash，错误日志会将堆栈信息打印出来

## 日志格式
```shell
yyyy/mm/dd hh:MM:ss.xxx xxx/../xx.go:86: [DEBUG] this is your file debug log.
```





