package util

import (
	"github.com/astaxie/beego"
	"os"
	"strings"
)

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func PathMkdir(path string) error {
	exist, err := PathExists(path)
	if err != nil {
		return err
	}
	if !exist {
		// 创建文件夹
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetFileExt(fileName string) string {
	if !CheckStrIsEmpty(fileName) && strings.LastIndex(fileName, ".") < len(fileName)-1 {
		return beego.Substr(fileName, strings.LastIndex(fileName, ".")+1, len(fileName))
	} else {
		return ""
	}
}
