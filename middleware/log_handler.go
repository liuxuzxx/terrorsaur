package middleware

import (
	"fmt"
	"github.com/kataras/iris/context"
)

func LogRequestInformationHandler(ctx context.Context) {
	fmt.Println("request信息内容是:", ctx.Host())
	ctx.Next()
}

func EncapsulationPage(ctx context.Context) {
	pageSize, sizeErr := ctx.URLParamInt("pageSize")
	pageNumber, numberErr := ctx.URLParamInt("pageNumber")

	if !(sizeErr != nil && numberErr != nil) {
		if sizeErr != nil {
			pageSize = 10
		}
		if numberErr != nil {
			pageNumber = 1
		}
		ctx.Params().SetImmutable("page", Page{PageNumber: pageNumber, PageSize: pageSize})
	}

	ctx.Next()
}

type Page struct {
	PageNumber int
	PageSize   int
}
