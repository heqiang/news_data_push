package util

import (
	"fmt"
	"testing"
)

func TestReadJson(t *testing.T) {
	newsList := ReadJsonFile("/Users/hq/GolandProjects/data_push/serivce/util/test.json")
	fmt.Println(newsList)
}
