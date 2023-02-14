package util

import (
	"fmt"
	"testing"
)

func TestReadJson(t *testing.T) {
	newsList := ReadJson("E:\\goproject\\data_push\\serivce\\util\\test.json")
	for _, k := range newsList {
		fmt.Println(k.URL)
	}
}
