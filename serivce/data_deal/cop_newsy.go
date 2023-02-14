package data_deal

import (
	"data_push/config"
	"data_push/serivce/util"
	"fmt"
	"go.uber.org/zap"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func CopyNews(sourceNewsPath, dstPath string, config *config.ServerConfig) {
	if !util.IsFileExist(sourceNewsPath) {
		zap.L().Error(fmt.Sprintf("新闻数据原始路径：%s 不存在", sourceNewsPath))
		return
	}
	imageDir := filepath.Join(dstPath, "image")
	util.MkAllDir(imageDir)

	fileDir := filepath.Join(dstPath, "file")
	util.MkAllDir(fileDir)

	zap.L().Info(fmt.Sprintf("开始复制%s第%s时的数据", config.PushDay, config.PushHour))

	fileList := util.ListDir(sourceNewsPath)
	if len(fileList) > 0 {
		newsCount := 0
		copyImageNum := 0
		copyFileNum := 0
		for _, dirPath := range fileList {
			currDatpath := filepath.Join(dirPath, config.PushDay)
			if util.IsFileExist(currDatpath) {
				for _, path := range util.ListDir(currDatpath) {
					_, fileName := filepath.Split(path)
					if strings.HasPrefix(fileName, fmt.Sprintf("%s%s", strings.ReplaceAll(config.PushDay, "-", ""), config.Host)) {
						if strings.HasSuffix(fileName, "newsty.json") {
							dstNewsPath := filepath.Join(dstPath, fmt.Sprintf("%s_newsty.json", fileName))
							newsCount += mergeJson(path, dstNewsPath)
						} else if strings.HasSuffix(fileName, "comments.json") {
							dstCommentPath := filepath.Join(dstPath, fmt.Sprintf("%s_comments.json", fileName))
							mergeJson(path, dstCommentPath)
						} else if strings.HasSuffix(fileName, "image") {
							copyImageNum += moveImageOrFile(path, imageDir)
						} else if strings.HasSuffix(fileName, "file") {
							copyFileNum += moveImageOrFile(path, fileDir)
						}

					}
				}
			}
		}
	}
}

func moveImageOrFile(srcPath, dstPath string) int {
	copyNum := 0
	for _, k := range util.ListDir(srcPath) {
		_ = os.Rename(filepath.Join(srcPath, k), filepath.Join(dstPath, k))
		copyNum++
	}
	return copyNum
}

func mergeJson(srcPathFile, dstPathFile string) int {
	mergeNum := 0

	file, _ := os.OpenFile(dstPathFile, os.O_APPEND, 0666)
	newsList := util.ReadJsonFile(srcPathFile)
	for _, k := range newsList {
		_, err := io.WriteString(file, k)
		if err != nil {
			return 0
		}
		mergeNum++
	}

	return mergeNum
}
