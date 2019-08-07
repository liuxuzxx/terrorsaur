package rest

import (
	"github.com/kataras/iris"
	"terrorsaur/common"
	"terrorsaur/service"
)

// 获取作者信息列表
// @Summary Author信息列表
// @Description 获取Author信息列表
// @Accept  json
// @Produce  json
// @Router /api/rattrap/ancient-article/author/{authorId:int}/information [get]
func AuthorInformation(context iris.Context) {
	authorId, _ := context.Params().GetInt("authorId")
	_, _ = context.JSON(common.Success(service.FetchAuthorByAuthorId(authorId)))
}
