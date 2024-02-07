package main

import (
	"flag"
	"fmt"
	"time"
	info "wfhome/Info"
	"wfhome/conf"
	"wfhome/file"
	"wfhome/logger"

	"github.com/rs/zerolog/log"
)

func devMode(level string) {
	// 设置日志级别
	logger.SetupLogger("debug")

	// 读取配置文件
	conf.LoadConf()

	// 备份日志文件路径
	backupPath := "test.log"

	// 读取日志文件
	file.ReadLog(backupPath)
}

func main() {
	level := flag.String("level", "info", "日志级别")
	flag.Parse()

	// 开发模式
	if *level == "dev" {
		devMode(*level)
		return
	}

	info.PrintInfo()

	// 设置日志级别
	logger.SetupLogger(*level)

	// 读取配置文件
	conf.LoadConf()

	// 开始时间戳
	star := time.Now().UnixMilli()

	// 备份日志文件路径
	backupPath := "logfile"

	// 备份日志文件
	file.CopyLog(backupPath)

	// 读取日志文件
	file.ReadLog(backupPath)

	// 当前时间戳
	now := time.Now().UnixMilli()

	// 计算运行时间
	log.Info().Int64("运行时间(单位: ms)", now-star).Msg("程序运行结束")

	// 等待用户输入
	fmt.Println("按回车键键退出")
	fmt.Scanln()
}
