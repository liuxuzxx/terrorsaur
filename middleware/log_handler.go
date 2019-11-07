package middleware

import (
	"fmt"
	"github.com/kataras/iris/v12/context"
)

func LogRequestInformationHandler(ctx context.Context) {
	fmt.Println("request信息内容是:", ctx.Host()+ctx.Path())
	ctx.Next()
}
