package result

import "terrorsaur/model"

//
//返回给前端的数据结果对象
//
type IdiomsResult struct {
	Id             int    `json:"id"`
	Term           string `json:"term"`
	Pronunciation  string `json:"pronunciation"`
	Interpretation string `json:"interpretation"`
	Source         string `json:"source"`
	Example        string `json:"example"`
}

func ConvertIdiomsDtoToResult(idioms model.Idioms) IdiomsResult {
	return IdiomsResult{
		Id:             idioms.Id,
		Term:           idioms.Term,
		Pronunciation:  idioms.Pronunciation,
		Interpretation: idioms.Interpretation,
		Source:         idioms.Source,
		Example:        idioms.Example,
	}
}

func ConvertIdiomsResults(idioms []model.Idioms) []IdiomsResult {
	idiomsResults := make([]IdiomsResult, len(idioms))
	for index, value := range idioms {
		idiomsResults[index] = ConvertIdiomsDtoToResult(value)
	}
	return idiomsResults
}
