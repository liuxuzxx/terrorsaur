package service

import (
	"terrorsaur/libs"
	"terrorsaur/model"
	"terrorsaur/result"
)

const TABLE_NAME = "article_type"

func FetchAllArticleType() []result.AncientArticleTypeResult {
	var articleTypes []model.ArticleType
	libs.Db.Table(TABLE_NAME).Find(&articleTypes)
	return result.ConvertTypeResults(articleTypes)
}

func FetchArticleTypeByTypeId(typeId int) result.AncientArticleTypeResult {
	var articleType model.ArticleType
	libs.Db.Table(TABLE_NAME).Where("type_id=?", typeId).First(&articleType)
	return result.ConvertTypeResult(articleType)
}
