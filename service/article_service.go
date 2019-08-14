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

func FetchArticlePageData(typeId, pageNumber int) {
}
