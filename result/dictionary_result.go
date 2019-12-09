package result

import "terrorsaur/model"

//
// @Date 2019-12-09 18:11-52
// @Author 刘旭
//主要放置dictionary有关的结果对象
//

type DictionaryTypeResult struct {
	Id             int    `json:"id"`
	DictionaryCode string `json:"dictionaryCode"`
	DictionaryName string `json:"dictionaryName"`
}

func ConvertDictionaryTypeToResult(dictionaryType model.DictionaryType) DictionaryTypeResult {
	return DictionaryTypeResult{
		Id:             dictionaryType.Id,
		DictionaryCode: dictionaryType.DictionaryCode,
		DictionaryName: dictionaryType.DictionaryName,
	}
}

func ConvertDictionaryTypeToResults(types []model.DictionaryType) []DictionaryTypeResult {
	typeResults := make([]DictionaryTypeResult, len(types))

	for index, value := range types {
		typeResults[index] = ConvertDictionaryTypeToResult(value)
	}
	return typeResults
}
