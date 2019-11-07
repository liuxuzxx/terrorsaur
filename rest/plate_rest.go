package rest

import (
	"github.com/kataras/iris/v12"
	"terrorsaur/common"
	"terrorsaur/service"
)

//
//板块的restful模块
//
// @Summary 全部板块数据信息
// @Description 全部板块数据信息
// @Accept  json
// @Produce  json
// @Router /api/rattrap/ancient-article/plate [get]
func AncientPlateInformation(ctx iris.Context) {
	ancientPlates := service.FetchAllPlates()
	_, _ = ctx.JSON(common.Success(ancientPlates))
}
