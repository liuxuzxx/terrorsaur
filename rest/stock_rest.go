package rest

import (
	"github.com/kataras/iris/v12"
	"terrorsaur/common"
	"terrorsaur/libs"
	"terrorsaur/service"
)

//
// @date   2020-04-13 14:22:12
// @author 刘旭
//
// 股票数据restful接口
//

var stockService = service.StockService{Db: libs.Db}

func FetchStockHistoryDatas(ctx iris.Context) {
	stockCode := ctx.Params().GetString("stockCode")
	_, _ = ctx.JSON(common.Success(stockService.LoadHistory(stockCode)))
}

func FetchStock(ctx iris.Context) {
	condition := ctx.URLParam("condition")
	_, _ = ctx.JSON(common.Success(stockService.FindStock(condition)))
}
