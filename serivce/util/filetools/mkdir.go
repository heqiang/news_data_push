package filetools

import (
	"fmt"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func MkAllDir(dirPath string) {
	err := os.MkdirAll(dirPath, 0666)
	if err != nil {
		zap.L().Error(fmt.Sprintf("目录：%s 创建失败", dirPath))
		return
	}
}

func MkFile(filePath string) {
	_, err := os.Create(filePath)
	if err != nil {
		zap.L().Error(fmt.Sprintf("文件 %s:创建失败", filePath))
		return
	}
}

func IsFileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	return true
}

func ListDir(dirPath string) []string {
	files, _ := ioutil.ReadDir(dirPath)
	fileOrDirList := make([]string, len(files))
	for _, f := range files {
		fileName := f.Name()
		if !strings.HasPrefix(fileName, ".") {
			filePath := filepath.Join(dirPath, fileName)
			fileOrDirList = append(fileOrDirList, filePath)
		}
		fmt.Println(f.Name())
	}
	return fileOrDirList
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func IsFile(path string) bool {
	return !IsDir(path)
}
