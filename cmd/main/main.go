package main

import (
	"data_push/config"
	"data_push/log_pkg"
	"data_push/serivce/data_deal"
	"fmt"
	"go.uber.org/zap"
	"os"
	"path/filepath"
)

func main() {
	config.InitConfig("dev")
	err := log_pkg.InitLogger(config.Config.LogConfig, "dev")
	if err != nil {
		panic(fmt.Sprintf("日志初始化错误,err:%s", err))
	}
	getwd, err := os.Getwd()
	if err != nil {
		zap.L().Error("获取当前路径失败")
		return
	}

	dstPath := filepath.Join(getwd, "data", config.Config.PushDay)
	// 复制当前小时的数据
	data_deal.CopyNews(config.Config.SourceNewsPath, dstPath, config.Config)

}
