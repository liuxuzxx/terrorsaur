package request

//
// @date   2020-03-28 15:49:33
// @author 刘旭
//
// 主要是放置有关于stock，也就是股票数据信息的请求对象
//

type Stock struct {
	Code string
	Name string
}

//来源于爬取财富网的股票信息对象
type StockJson struct {
	Rc   int64         `json:"rc"`
	Rt   int64         `json:"rt"`
	Svr  int64         `json:"svr"`
	Lt   int64         `json:"lt"`
	Full int64         `json:"full"`
	Data StockJsonData `json:"data"`
}

type StockJsonData struct {
	Total int64       `json:"total"`
	Diff  []StockDiff `json:"diff"`
}

type StockDiff struct {
	F12 string `json:"f12"`
	F14 string `json:"f14"`
}
