package service

import (
	"terrorsaur/libs"
	"terrorsaur/model"
	"terrorsaur/result"
)

//
// @Date 2019-12-09 18:17-37
// @Author 刘旭
//
//字典操作服务
//

const (
	DictionaryTypeTableName string = "dictionary_type"
)

func FetchAllDictionaryTypes() []result.DictionaryTypeResult {
	var dictionaryTypes []model.DictionaryType
	libs.Db.Table(DictionaryTypeTableName).Find(&dictionaryTypes)
	return result.ConvertDictionaryTypeToResults(dictionaryTypes)
}
