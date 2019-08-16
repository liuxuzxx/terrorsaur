package service

import (
	"terrorsaur/common"
	"terrorsaur/libs"
	"terrorsaur/model"
	"terrorsaur/result"
)

const ArticleTypeTableName = "article_type"
const ArticleTableName = "article"
const ArticleAttributeTableName = "article_attribute"

func FetchAllArticleType() []result.AncientArticleTypeResult {
	var articleTypes []model.ArticleType
	libs.Db.Table(ArticleTypeTableName).Find(&articleTypes)
	return result.ConvertTypeResults(articleTypes)
}

func FetchArticleTypeByTypeId(typeId int) result.AncientArticleTypeResult {
	var articleType model.ArticleType
	libs.Db.Table(ArticleTypeTableName).Where("type_id=?", typeId).First(&articleType)
	return result.ConvertTypeResult(articleType)
}

func FetchArticlePageData(typeId int, page common.Page) []result.ArticleResult {
	var articles []model.Article
	libs.Db.Table(ArticleTableName).Where("article_attribute.attribute_code=1 and article_attribute.attribute_value=?", typeId).Select("article.article_id,article.title").Joins("left join article_attribute on article.article_id = article_attribute.article_id").Limit(page.PageSize).Offset((page.PageNumber - 1) * page.PageSize).Find(&articles)
	return result.ConvertArticleResults(articles)
}

func FetchArticleDetail(articleId int) result.ArticleResult {
	var article model.Article
	libs.Db.Table(ArticleTableName).Where("article_id=?", articleId).First(&article)
	articleResult := result.ConvertArticleResult(article)
	var attributes []model.ArticleAttribute
	libs.Db.Table(ArticleAttributeTableName).Where("article_id=?", articleId).Find(&attributes)
	attributeResults := result.ConvertArticleAttributeResults(attributes)
	articleResult.Attributes = attributeResults
	return articleResult
}
