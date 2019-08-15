package result

import (
	"terrorsaur/model"
	"time"
)

type AncientArticleTypeResult struct {
	TypeId    int    `json:"typeId"`
	TypeName  string `json:"typeName"`
	Detail    string `json:"detail"`
	TypeOrder int    `json:"typeOrder"`
}

type ArticleResult struct {
	ArticleId  int       `json:"articleId"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Comment    string    `json:"comment"`
	CreateBy   string    `json:"createBy"`
	CreateTime time.Time `json:"createTime"`
	UpdateBy   string    `json:"updateBy"`
	UpdateTime time.Time `json:"updateTime"`
}

type ArticleAttributeResult struct {
	ArticleId         int    `json:"articleId"`
	AttributeCode     int    `json:"attributeCode"`
	AttributeValue    string `json:"attributeValue"`
	AttributeCodeStr  string `json:"attributeCodeStr"`
	AttributeValueStr string `json:"attributeValueStr"`
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

func ConvertArticleResults(articles []model.Article) []ArticleResult {
	articleResults := make([]ArticleResult, len(articles))
	for index, value := range articles {
		articleResults[index] = ConvertArticleResult(value)
	}
	return articleResults
}

func ConvertArticleResult(article model.Article) ArticleResult {
	return ArticleResult{
		ArticleId:  article.ArticleId,
		Title:      article.Title,
		Content:    article.Content,
		Comment:    article.Comment,
		CreateBy:   article.CreateBy,
		CreateTime: article.CreateTime,
		UpdateBy:   article.UpdateBy,
		UpdateTime: article.UpdateTime,
	}
}

func ConvertArticleAttributeResult(attribute model.ArticleAttribute) ArticleAttributeResult {
	return ArticleAttributeResult{
		ArticleId:      attribute.ArticleId,
		AttributeCode:  attribute.AttributeCode,
		AttributeValue: attribute.AttributeValue,
	}
}
