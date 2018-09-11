package util

import (
	"bufio"
	"os"
	"strings"
)

const (
	FN = "./conf/app.conf"
)

//str配置文件中的key
func Get(str string) string {
	if len(str) < 1 {
		return ""
	}
	f, err := os.Open(FN)
	defer f.Close()
	if err != nil {
		return ""
	}
	buf := bufio.NewReader(f)
	for {
		line, _ := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		strs := strings.Split(line, "=")
		if strings.TrimSpace(strs[0]) == str {
			return strings.TrimSpace(strs[1])
		}
	}
	return ""
}

func GetOrDefault(key string, defaultValue string) string {
	value := Get(key)
	if len(value) > 0 {
		return value
	} else {
		return defaultValue
	}
}
