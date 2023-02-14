package config

import (
	"data_push/serivce/util"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/golang-module/carbon"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"strconv"
)

func InitConfig(env string) {
	rootPath, err := os.Getwd()
	if err != nil {
		msg := fmt.Sprintf("根目录路径获取错误:%s", err)
		zap.L().Error(msg)
		return
	}

	configFilePrefix := "config"
	configFileName := fmt.Sprintf("%s_pro.yaml", configFilePrefix)
	if env == "dev" {
		configFileName = fmt.Sprintf("%s_dev.yaml", configFilePrefix)
	}

	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(&Config); err != nil {
		panic(err)
	}

	Config.VideoPath = filepath.Join(rootPath, "news_video")
	util.MkAllDir(Config.VideoPath)
	//是否补推
	isSupp := v.Get("is_supplementary_push")
	if !isSupp.(bool) {
		Config.PushHour = strconv.Itoa(carbon.Now().Hour())
		Config.PushDay = carbon.Now().ToDateString()
	}
	Config.DataPath = filepath.Join(rootPath, fmt.Sprintf("news-%s", Config.PushHour))
	Config.ZipPath = filepath.Join(rootPath, "54_data")
	util.MkAllDir(Config.ZipPath)
	Config.JsonSavePath = filepath.Join(rootPath, "hours_data")
	util.MkAllDir(Config.JsonSavePath)

	//配置热加载
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		if err := viper.Unmarshal(&Config); err != nil {
			panic(fmt.Errorf("unmarshal conf failed err:%s\n", err))
		}
	})
	return
}
