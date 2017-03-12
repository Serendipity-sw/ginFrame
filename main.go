/**
场景导航服务程序入口
创建人:邵炜
创建时间:2017年2月8日17:56:01
*/
package main

import (
	"./common"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/guotie/config"
	"github.com/guotie/deferinit"
	"github.com/smtc/glog"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

var (
	pidStrPath          = "./sceneMarketingService.pid"
	configFn            = flag.String("config", "./config.json", "config file path")
	debugFlag           = flag.Bool("d", false, "debug mode")
	serverListeningPort int         //服务监听端口
	logsDir             string      //程序记录日志存放目录
	rootPrefix          string      //服务运行所需要的二级目录名称
	loadTemplates       string      //需要加载运行的html模板路径
	rt                  *gin.Engine //web初始化后变量
	serverPidFilePath   string      //程序运行PID文件目录
)

/**
初始化日志类
创建人:邵炜
创建时间:2017年2月8日18:26:29
输入参数:是否为调试模式(true为调试模式 false为正式运行模式)
*/
func logInit(debug bool) {
	option := map[string]interface{}{
		"typ": "file",
	}
	if len(logsDir) != 0 {
		option["dir"] = logsDir
	}
	glog.InitLogger(common.If(debug, glog.DEV, glog.PRO).(int), option)
}

/**
初始化web工具
创建人:邵炜
创建时间:2017年2月9日13:45:26
输入参数:是否为调试模式(true为调试模式 false为正式运行模式)
*/
func ginInit(debug bool) {
	//设置gin的工作方式
	gin.SetMode(common.If(debug, gin.DebugMode, gin.ReleaseMode).(string))
	rt = gin.Default()
	if len(loadTemplates) != 0 {
		rt.LoadHTMLGlob(fmt.Sprintf("%s/*", loadTemplates))
	}
	setGinRouter(rt)
	go rt.Run(fmt.Sprintf(":%d", serverListeningPort))
}

/**
服务运行
创建人:邵炜
创建时间:2017年2月8日18:01:18
输入参数:配置文件路径 是否为调试模式(d表示为调试模式,否则为正式运行模式)
*/
func serverRun(cfn string, debug bool) {
	config.ReadCfg(cfn)
	configRead()
	logInit(debug)

	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("set many cpu successfully!")

	deferinit.InitAll()
	fmt.Println("init all module successfully!")

	deferinit.RunRoutines()
	fmt.Println("init all run successfully!")

	ginInit(debug)
	fmt.Println("ginInit run successfully!")
}

/**
服务停止
创建人:邵炜
创建时间:2017年2月9日14:06:27
*/
func serverExit() {
	deferinit.StopRoutines()
	fmt.Println("stop routine successfully!")

	deferinit.FiniAll()
	fmt.Println("stop all modules successfully!")

	glog.Close()
}

/**
服务构造函数(程序启动主入口)
创建人:邵炜
创建时间:2017年2月9日14:08:21
*/
func main() {
	if common.CheckPid(pidStrPath) {
		return
	}
	flag.Parse()
	serverRun(*configFn, *debugFlag)
	c := make(chan os.Signal, 1)
	common.WritePid(pidStrPath)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	//信号等待
	<-c
	common.RmPidFile(pidStrPath)
	serverExit()
	os.Exit(0)
}
