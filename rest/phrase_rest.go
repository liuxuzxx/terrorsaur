package rest

import (
	"github.com/kataras/iris/v12"
	"terrorsaur/common"
	"terrorsaur/service"
)

//
//成语的RESTFul接口
//
// @Summary 获取成语的分页数据结果对象
// @Description 就是获取成语的分页数据对象
// @Accept  json
// @Produce  json
// @Router /api/rattrap/ancient-article/phrase/information [get]
func IdiomsInformation(ctx iris.Context) {
	page := common.ConvertToPage(ctx)
	idiomsResults := service.FetchIdiomsPageData(page)
	_, _ = ctx.JSON(common.Success(idiomsResults))
}
