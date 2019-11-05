package service

import (
	"terrorsaur/common"
	"terrorsaur/libs"
	"terrorsaur/model"
	"terrorsaur/result"
)

const (
	IdiomsTableName string = "idioms"
)

func FetchIdiomsPageData(page common.Page) []result.IdiomsResult {
	var idioms []model.Idioms
	libs.Db.Table(IdiomsTableName).Select("id,term,pronunciation,interpretation,source,example").Limit(page.PageSize).Offset((page.PageNumber - 1) * page.PageSize).Find(&idioms)
	return result.ConvertIdiomsResults(idioms)
}
