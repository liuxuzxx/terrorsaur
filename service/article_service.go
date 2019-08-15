package service

import (
	"terrorsaur/libs"
	"terrorsaur/model"
	"terrorsaur/result"
)

const ArticleTypeTableName = "article_type"
const ArticleTableName = "article"

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

func FetchArticlePageData(typeId int) []result.ArticleResult {
	var articles []model.Article
	libs.Db.Table(ArticleTableName).Where("article_attribute.attribute_code=1 and article_attribute.attribute_value=?", typeId).Select("article.article_id,article.title").Joins("left join article_attribute on article.article_id = article_attribute.article_id").Limit(10).Offset(1).Find(&articles)
	return result.ConvertArticleResults(articles)
}
