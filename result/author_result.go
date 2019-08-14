package result

import "terrorsaur/model"

type AuthorResult struct {
	AuthorId int    `json:"authorId"`
	Name     string `json:"name"`
	Dynasty  string `json:"dynasty"`
	Detail   string `json:"detail"`
}

func ConvertAuthorToResult(author model.Author) AuthorResult {
	return AuthorResult{
		author.AuthorId,
		author.Name,
		author.Dynasty,
		author.Detail,
	}
}
