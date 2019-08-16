package common

import (
	"github.com/kataras/iris"
)

type Page struct {
	PageSize   int
	PageNumber int
}

func ConvertToPage(ctx iris.Context) Page {
	pageSize, sizeErr := ctx.URLParamInt("pageSize")
	pageNumber, numberErr := ctx.URLParamInt("pageNumber")

	if !(sizeErr != nil && numberErr != nil) {
		if sizeErr != nil {
			pageSize = 10
		}
		if numberErr != nil {
			pageNumber = 1
		}
	}

	return Page{
		PageSize:   pageSize,
		PageNumber: pageNumber,
	}
}
