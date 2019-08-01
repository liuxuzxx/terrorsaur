package rest

import (
	"github.com/kataras/iris"
	"terrorsaur/common"
	"terrorsaur/service"
)

func ArticleTypePageData(context iris.Context) {
	articleTypes := service.FetchAllArticleType()
	_, _ = context.JSON(common.Success(articleTypes))
}

func ArticleTypeInformation(context iris.Context) {
	typeId, _ := context.Params().GetInt("typeId")
	articleType := service.FetchArticleTypeByTypeId(typeId)
	_, _ = context.JSON(common.Success(articleType))
}
