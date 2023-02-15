package data_deal

import (
	"data_push/config"
	"data_push/serivce/redispkg"
	"data_push/serivce/util"
	"data_push/serivce/util/filetools"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"path/filepath"
)

func DupByRedis(dstPath string) {
	dupJsonFie := filepath.Join(dstPath, fmt.Sprintf("%s_filter.json", config.Config.DayHours))
	if !filetools.IsFileExist(dupJsonFie) {
		zap.L().Error(fmt.Sprintf("文件:%s 不存在", dupJsonFie))
		return
	}
	zap.L().Info(fmt.Sprintf("开始对文件:%s 去重", dupJsonFie))
	file, err := os.OpenFile(dupJsonFie, os.O_RDONLY, 0666)
	defer file.Close()
	if err != nil {
		zap.L().Error(fmt.Sprintf("文件%s打开失败", dupJsonFie))
		return
	}
	var newsList []map[string]interface{}
	fileBytes, err := ioutil.ReadAll(file)
	dupCount := 0
	gjson.ForEachLine(string(fileBytes), func(line gjson.Result) bool {
		var newsLine map[string]interface{}
		err := json.Unmarshal([]byte(line.String()), &newsLine)
		if err != nil {
			zap.L().Warn(fmt.Sprintf("去重阶段 %s文件 json反序列化失败", dstPath))
			return false
		}
		if redispkg.Add(newsLine["news_id"].(string)) {
			newsList = append(newsList, newsLine)
		} else {
			dupCount++
			zap.L().Info(fmt.Sprintf("去重 去掉第%d条数据", dupCount))
		}
		return true
	})

	util.WriteJson(newsList, dstPath, fmt.Sprintf("%s_json", config.Config.DayHours))
	err = os.Remove(dupJsonFie)
	if err != nil {
		zap.L().Info(fmt.Sprintf("_filter.json文件 %s删除失败", dupJsonFie))
		return
	}
	zap.L().Info(fmt.Sprintf("去重完毕,一共去除%d条数据", dupCount))

}
