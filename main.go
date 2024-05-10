package main

import "log"

// init用于初始化操作，这里是操作日志的输出格式，log.Ltime表示在日志中输出当前时间，log.Lshortfile表示在日志中输出调用日志语句的文件名和行号
func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
}

func initSettings() {
	setting.AllInit()
}

func main() {
	initSettings()
}
