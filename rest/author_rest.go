package rest

import (
	"github.com/kataras/iris/v12"
	"terrorsaur/common"
	"terrorsaur/service"
)

// @Summary Author信息列表
// @Description 获取Author信息列表
// @Accept  json
// @Produce  json
// @Router /api/rattrap/ancient-article/author/{authorId}/information [get]
func AuthorInformation(context iris.Context) {
	authorId, _ := context.Params().GetInt("authorId")
	_, _ = context.JSON(common.Success(service.FetchAuthorByAuthorId(authorId)))
}
