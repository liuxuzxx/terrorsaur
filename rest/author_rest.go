package rest

import (
	"github.com/kataras/iris"
	"terrorsaur/common"
	"terrorsaur/service"
)

func AuthorInformation(context iris.Context) {
	authorId, _ := context.Params().GetInt("authorId")
	_, _ = context.JSON(common.Success(service.FetchAuthorByAuthorId(authorId)))
}
