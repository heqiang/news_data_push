package util

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
	"io"
	"os"
)

func ReadJsonFile(jsonPath string) []string {
	srcFile, err := os.OpenFile(jsonPath, os.O_RDONLY, 0666)
	defer srcFile.Close()
	if err != nil {
		zap.L().Warn(fmt.Sprintf("文件%s打开失败", jsonPath))
		return nil
	}
	fileBytes, err := io.ReadAll(srcFile)
	if err != nil {
		zap.L().Warn(fmt.Sprintf("文件%s读取失败", jsonPath))
		return nil
	}
	var newsList []string

	gjson.ForEachLine(string(fileBytes), func(line gjson.Result) bool {
		newsList = append(newsList, line.String())
		return true
	})

	return newsList
}

func WriteJson(jsonData []map[string]interface{}, jsonFilePath string) {
	dstFile, err := os.OpenFile(jsonFilePath, os.O_APPEND, 0666)
	if err != nil {
		zap.L().Warn("新闻json写入失败")
		return
	}

	for _, v := range jsonData {
		vBytes, err := json.Marshal(v)
		if err != nil {
			return
		}
		_, err = io.WriteString(dstFile, string(vBytes))
		if err != nil {
			return
		}
	}

}
