package common

import (
	"github.com/kataras/iris/v12"
)

type Page struct {
	PageSize   int
	PageNumber int
}

func ConvertToPage(ctx iris.Context) Page {
	pageSize, sizeErr := ctx.URLParamInt("pageSize")
	pageNumber, numberErr := ctx.URLParamInt("pageNumber")

	if sizeErr != nil || pageSize == -1 {
		pageSize = 10
	}
	if numberErr != nil || pageNumber == -1 {
		pageNumber = 1
	}

	return Page{
		PageSize:   pageSize,
		PageNumber: pageNumber,
	}
}
