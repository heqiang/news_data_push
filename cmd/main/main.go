package main

import (
	"data_push/config"
	"data_push/log_pkg"
	"data_push/serivce/data_deal"
	"data_push/serivce/redispkg"
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

	err = redispkg.InitRedis(config.Config.RedisConfig)
	if err != nil {
		panic(fmt.Sprintf("redis db 初始化错误,err:%s", err))
	}

	getwd, err := os.Getwd()
	if err != nil {
		zap.L().Error("获取当前路径失败")
		return
	}

	dstPath := filepath.Join(getwd, "data", config.Config.PushDay)
	// 复制当前小时的数据
	data_deal.CopyNews(config.Config.SourceNewsPath, dstPath, config.Config)
	data_deal.FilterNews(dstPath)

	data_deal.DupByRedis(dstPath)

	// 数据写入csv
}
