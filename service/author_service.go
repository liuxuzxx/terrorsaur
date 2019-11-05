package service

import (
	"terrorsaur/libs"
	"terrorsaur/model"
	"terrorsaur/result"
)

const (
	AuthorTableName string = "author"
)

func FetchAuthorByAuthorId(authorId int) result.AuthorResult {
	var author model.Author
	libs.Db.Table(AuthorTableName).Where("author_id=?", authorId).First(&author)
	return result.ConvertAuthorToResult(author)
}
