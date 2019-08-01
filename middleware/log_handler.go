package middleware

import (
	"fmt"
	"github.com/kataras/iris/context"
)

func LogRequestInformationHandler(ctx context.Context) {
	fmt.Println("request信息内容是:", ctx.Host())
	ctx.Next()
}
