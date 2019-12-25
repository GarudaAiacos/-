package main

import (
	"strings"
)

func replaceSpace(str string) string {
	//-1参数表示替换次数，负数表示没有限制
	str = strings.Replace(str, " ", "%20", -1)
	return str
}
