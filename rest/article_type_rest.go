package rest

import (
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
	page := common.ConvertToPage(ctx)
	articleResults := service.FetchArticlePageData(typeId, page)
	_, _ = ctx.JSON(common.Success(articleResults))
}

func ArticleDetailInformation(ctx iris.Context) {
	articleId, _ := ctx.Params().GetInt("articleId")
	detail := service.FetchArticleDetail(articleId)
	_, _ = ctx.JSON(common.Success(detail))
}
