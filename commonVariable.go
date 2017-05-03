package main

import (
	"flag"
	"github.com/gin-gonic/gin"
)

var (
	pidStrPath          = "./run-program.pid"
	configFn            = flag.String("config", "./config.json", "config file path")
	debugFlag           = flag.Bool("d", false, "debug mode")
	serverListeningPort int         //服务监听端口
	logsDir             string      //程序记录日志存放目录
	rootPrefix          string      //服务运行所需要的二级目录名称
	loadTemplates       string      //需要加载运行的html模板路径
	rt                  *gin.Engine //web初始化后变量
	fileProgram         string      //程序生成文件存放目录
)
