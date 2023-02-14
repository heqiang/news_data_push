package util

import (
	"os"
	"path/filepath"
	"testing"
)

func TestMkAllDir(t *testing.T) {
	getwd, err := os.Getwd()
	if err != nil {
		return
	}
	MkAllDir(filepath.Join(getwd, "54_data/test"))
}

func TestMkFile(t *testing.T) {
	MkFile("test.txt")
}

func TestListDir(t *testing.T) {
	ListDir("E:\\goproject\\data_push")
}
