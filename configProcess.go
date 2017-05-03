package main

import (
	"fmt"
	"github.com/guotie/config"
	"github.com/swgloomy/gutil"
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
	err := gutil.CreateFileProcess(logsDir)
	if err != nil && len(logsDir) != 0 {
		fmt.Printf("configRead CreateFileProcess run err! logDir: %s err: %s \n", logsDir, err.Error())
		os.Exit(1)
	}
	loadTemplates = strings.TrimSpace(config.GetStringDefault("loadTemplates", ""))
	err = gutil.CreateFileProcess(loadTemplates)
	if err != nil && len(loadTemplates) != 0 {
		fmt.Printf("configRead CreateFileProcess run err! loadTemplates: %s err: %s \n", loadTemplates, err.Error())
		os.Exit(1)
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
	fileProgram = strings.TrimSpace(config.GetStringMust("fileProgram"))
	err = gutil.CreateFileProcess(fileProgram)
	if err != nil && len(fileProgram) != 0 {
		fmt.Printf("configRead CreateFileProcess run err! fileProgram: %s err: %s \n", fileProgram, err.Error())
		os.Exit(1)
	}
}
