package result

import (
	"terrorsaur/model"
)

type AncientArticleTypeResult struct {
	TypeId    int    `json:"typeId"`
	TypeName  string `json:"typeName"`
	Detail    string `json:"detail"`
	TypeOrder int    `json:"typeOrder"`
}

func ConvertTypeResults(articleTypes []model.ArticleType) []AncientArticleTypeResult {
	typeResults := make([]AncientArticleTypeResult, len(articleTypes))
	for index, value := range articleTypes {
		typeResults[index] = ConvertTypeResult(value)
	}
	return typeResults
}

func ConvertTypeResult(articleType model.ArticleType) AncientArticleTypeResult {
	return AncientArticleTypeResult{
		TypeId:    articleType.TypeId,
		TypeName:  articleType.TypeName,
		Detail:    articleType.Detail,
		TypeOrder: articleType.TypeOrder,
	}
}
