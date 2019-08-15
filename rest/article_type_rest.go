package rest

import (
	"fmt"
	"github.com/kataras/iris"
	"terrorsaur/common"
	"terrorsaur/service"
)

func ArticleTypePageData(ctx iris.Context) {
	articleTypes := service.FetchAllArticleType()
	_, _ = ctx.JSON(common.Success(articleTypes))
}

func ArticleTypeInformation(ctx iris.Context) {
	typeId, _ := ctx.Params().GetInt("typeId")
	articleType := service.FetchArticleTypeByTypeId(typeId)
	_, _ = ctx.JSON(common.Success(articleType))
}

func ArticlesInformation(ctx iris.Context) {
	typeId, _ := ctx.Params().GetInt("typeId")
	page := ctx.Params().Get("page")
	fmt.Println("获取到的数据:", typeId, page)
	articleResults := service.FetchArticlePageData(typeId)
	_, _ = ctx.JSON(common.Success(articleResults))
}
