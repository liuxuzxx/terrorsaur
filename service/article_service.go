package service

import (
	"terrorsaur/libs"
	"terrorsaur/model"
	"terrorsaur/result"
)

const TableName = "article_type"

func FetchAllArticleType() []result.AncientArticleTypeResult {
	var articleTypes []model.ArticleType
	libs.Db.Table(TableName).Find(&articleTypes)
	return result.ConvertTypeResults(articleTypes)
}

func FetchArticleTypeByTypeId(typeId int) result.AncientArticleTypeResult {
	var articleType model.ArticleType
	libs.Db.Table(TableName).Where("type_id=?", typeId).First(&articleType)
	return result.ConvertTypeResult(articleType)
}
