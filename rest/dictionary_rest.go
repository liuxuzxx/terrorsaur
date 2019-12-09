package rest

import (
	"github.com/kataras/iris/v12"
	"terrorsaur/common"
	"terrorsaur/service"
)

//
// @Date 2019-12-09 18:23-32
// @Author 刘旭
//
//字典的restful接口
//
func DictionaryTypeInformation(ctx iris.Context) {
	dictionaryTypes := service.FetchAllDictionaryTypes()
	_, _ = ctx.JSON(common.Success(dictionaryTypes))
}
