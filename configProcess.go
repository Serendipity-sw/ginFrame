package main

import (
	"./common"
	"fmt"
	"github.com/guotie/config"
	"os"
	"strings"
)

/**
配置文件读取
创建人:邵炜
创建时间:2017年2月8日18:24:34
*/
func configRead() {
	logsDir = strings.TrimSpace(config.GetStringDefault("logDir", ""))
	err := common.CreateFileProcess(logsDir)
	if err != nil && len(logsDir) != 0 {
		fmt.Printf("configRead CreateFileProcess run err! logDir: %s err: %s \n", logsDir, err.Error())
		os.Exit(1)
	}
	loadTemplates = strings.TrimSpace(config.GetStringDefault("loadTemplates", ""))
	err = common.CreateFileProcess(loadTemplates)
	if err != nil && len(loadTemplates) != 0 {
		fmt.Printf("configRead CreateFileProcess run err! loadTemplates: %s err: %s \n", loadTemplates, err.Error())
		os.Exit(1)
	}
	serverPidFilePath = strings.TrimSpace(config.GetStringMust("serverPidFilePath"))
	if len(serverPidFilePath) == 0 {
		fmt.Println("configRead pid key read err!")
	}
	serverListeningPort = config.GetIntDefault("serverListeningPort", 8000)
	rootPrefix = strings.TrimSpace(config.GetStringDefault("rootPrefix", ""))
	if len(rootPrefix) != 0 {
		if !strings.HasPrefix(rootPrefix, "/") {
			rootPrefix = "/" + rootPrefix
		}
		if strings.HasSuffix(rootPrefix, "/") {
			rootPrefix = rootPrefix[0 : len(rootPrefix)-1]
		}
	}
}
