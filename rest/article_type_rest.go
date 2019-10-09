package rest

import (
	"github.com/kataras/iris"
	"terrorsaur/common"
	"terrorsaur/service"
)

// @Summary 全部文章类型的数据
// @Description 全部文章类型的数据信息
// @Accept  json
// @Produce  json
// @Router /api/rattrap/ancient-article/article-type/types [get]
func ArticleTypePageData(ctx iris.Context) {
	articleTypes := service.FetchAllArticleType()
	_, _ = ctx.JSON(common.Success(articleTypes))
}

// @Summary 根据类型ID获取文章类型的详细信息
// @Description 根据文章的类型ID(typeId),获取文章类型的详细信息
// @Accept  json
// @Produce  json
// @Router /api/rattrap/ancient-article/article-type/{typeId}/information [get]
func ArticleTypeInformation(ctx iris.Context) {
	typeId, _ := ctx.Params().GetInt("typeId")
	articleType := service.FetchArticleTypeByTypeId(typeId)
	_, _ = ctx.JSON(common.Success(articleType))
}

// @Summary 根据类型获取文章的分页数据信息
// @Description 根据类型typeId，获取该类型文章的分页数据信息
// @Accept  json
// @Produce  json
// @Router /api/rattrap/ancient-article/article/{typeId} [get]
func ArticlesInformation(ctx iris.Context) {
	typeId, _ := ctx.Params().GetInt("typeId")
	page := common.ConvertToPage(ctx)
	articleResults := service.FetchArticlePageData(typeId, page)
	_, _ = ctx.JSON(common.Success(articleResults))
}

// @Summary 获取文章的详细信息
// @Description 获取文章的详细信息,根据文章ID
// @Accept  json
// @Produce  json
// @Router /api/rattrap/ancient-article/article/{articleId}/detail [get]
func ArticleDetailInformation(ctx iris.Context) {
	articleId, _ := ctx.Params().GetInt("articleId")
	detail := service.FetchArticleDetail(articleId)
	_, _ = ctx.JSON(common.Success(detail))
}
