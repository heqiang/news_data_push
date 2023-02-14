package data_deal

import (
	"data_push/serivce/util"
	"encoding/json"
	"fmt"
	mapset "github.com/deckarep/golang-set"
	"go.uber.org/zap"
	"os"
	"strings"
	"time"
)

var dropSites = []interface{}{"U港生活", "论尽媒体", "新加坡联合早报网", "灼见名家", "Yam番薯藤", "巨亨网新闻", "香港电台网站",
	"台湾旅行趣", "台湾产经新闻网", "中华新知传媒", "澎湖县政府入口网"}

var dirtyData = []string{"JavaScript", "APP看新聞", "browser", "以上瀏覽", "新聞關鍵字：", "Firefox"}

func FilterNews(dstPath string) {
	dropNewsCount := 0
	removeNewsName := ""

	dataList := util.ListDir(dstPath)
	if len(dataList) > 0 {
		allNews := make([]map[string]interface{}, 0)
		for _, v := range dataList {
			if strings.HasSuffix(v, "newsty.json") {
				newsList := util.ReadJsonFile(v)
				removeNewsName = v
				for _, v := range newsList {
					jsonData := make(map[string]interface{})
					err := json.Unmarshal([]byte(v), &jsonData)
					if err != nil {
						dropNewsCount++
						zap.L().Error(fmt.Sprintf("json反序列化失败,丢失第%d条数据", dropNewsCount))
						continue
					}
					jsonData["ori_news_id"] = jsonData["news_id"]

					if mapset.NewSetFromSlice(dropSites).Contains(jsonData["source_name"].(string)) {
						dropNewsCount++
						zap.L().Error(fmt.Sprintf("无关网站,丢失第%d条数据", dropNewsCount))
						continue
					}

					if jsonData["time"] == "" {
						jsonData["time"] = "9999-01-01 00:00:00"
					}
					if jsonData["time"] != "9999-01-01 00:00:00" {
						_, err = time.Parse("2006-01-02 15:04:05", jsonData["time"].(string))
						if err != nil {
							zap.L().Warn(fmt.Sprintf("时间 %s 解析失败,丢失第%d条数据,新闻id:%s", jsonData["time"], dropNewsCount, jsonData["news_id"]))
							dropNewsCount++
							continue
						}
					}
					content := jsonData["content"].([]map[string]interface{})
					if len(content) == 0 {
						dropNewsCount++
						zap.L().Warn(fmt.Sprintf("内容缺失,丢失第%d条数据,新闻id:%s", dropNewsCount, jsonData["news_id"]))
						continue
					}
					newContent := make([]map[string]interface{}, len(content))
					for _, v := range content {
						if v["type"] == "text" {
							conData := v["data"].(string)
							if len(conData) != 0 {
								if !filterDirty(dropSites, conData) {
									newContent = append(newContent, v)
								}

							} else {
								dropNewsCount++
								zap.L().Warn(fmt.Sprintf("正文内容有误,丢失第%d条数据,新闻id:%s", dropNewsCount, jsonData["news_id"]))
							}
						} else if v["type"] == "image" {

						}
					}

					jsonData["special_name"] = ""
					jsonData["special_keyword"] = ""
					allNews = append(allNews, jsonData)
				}
			}
		}
		util.WriteJson(allNews, dstPath)
	}
	if removeNewsName != "" {
		//删除json文件
		err := os.Remove(removeNewsName)
		if err != nil {
			zap.L().Warn(fmt.Sprintf("文件 %s 删除失败", removeNewsName))
			return
		}
	} else {
		zap.L().Warn("当前小时包下没有新闻的json数据")
	}

}

func filterDirty(dirtyData []interface{}, substr string) bool {
	res := false
	for _, v := range dirtyData {
		if strings.Contains(substr, v.(string)) {
			res = true
		}
	}
	return res

}
