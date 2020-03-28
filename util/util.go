package util

//
// @date   2020-03-28 17:18:09
// @author 刘旭
//
// 主要是放置一些工具函数
//

import (
	"github.com/axgle/mahonia"
)

func ConvertToString(source string, sourceCode string, targetCode string) string {
	newSourceCode := mahonia.NewDecoder(sourceCode)
	result := newSourceCode.ConvertString(source)
	newTargetCode := mahonia.NewDecoder(targetCode)
	_, cdata, _ := newTargetCode.Translate([]byte(result), true)
	return string(cdata)
}

func GBKToUTF8(source string) string {
	return ConvertToString(source, "gbk", "utf-8")
}
