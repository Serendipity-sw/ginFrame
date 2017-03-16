/**
日志处理类
创建人:邵炜
创建时间:2017年3月13日14:26:47
*/
package main

import (
	"github.com/smtc/glog"
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
	if debug {
		glog.InitLogger(glog.DEV,option)
	}else{
		glog.InitLogger(glog.PRO,option)
	}
}
