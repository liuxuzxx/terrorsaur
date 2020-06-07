package result

import (
	"terrorsaur/model"
	"time"
)

//
// @date   2020-04-13 14:08:21
// @author 刘旭
//
// stock股票信息结果对象
//

type StockHistoryDataResult struct {
	HistoryId                 int64     `json:"historyId"`
	StockCode                 string    `json:"stockCode"`
	DateTime                  time.Time `json:"dateTime"`
	StartPrice                float64   `json:"startPrice"`
	EndPrice                  float64   `json:"endPrice"`
	YesterdayEndPrice         float64   `json:"yesterdayEndPrice"`
	MaxPrice                  float64   `json:"maxPrice"`
	MinPrice                  float64   `json:"minPrice"`
	RiseFall                  float64   `json:"riseFall"`
	TurnoverRate              float64   `json:"turnoverRate"`
	Volume                    int64     `json:"volume"`
	Turnover                  float64   `json:"turnover"`
	TotalMarketCapitalization float64   `json:"totalMarketCapitalization"`
	MarketCapitalization      float64   `json:"marketCapitalization"`
}

type StockResult struct {
	StockId   int64  `json:"stockId"`
	StockCode string `json:"stockCode"`
	StockName string `json:"stockName"`
	Type      string `json:"type"`
	StartDate string `json:"startDate"`
}

func ConvertToResult(historyData model.StockHistoryData) StockHistoryDataResult {
	return StockHistoryDataResult{
		HistoryId:                 historyData.HistoryId,
		StockCode:                 historyData.StockCode,
		DateTime:                  historyData.DateTime,
		StartPrice:                historyData.StartPrice,
		EndPrice:                  historyData.EndPrice,
		YesterdayEndPrice:         historyData.YesterdayEndPrice,
		MaxPrice:                  historyData.MaxPrice,
		MinPrice:                  historyData.MinPrice,
		RiseFall:                  historyData.RiseFall,
		TurnoverRate:              historyData.TurnoverRate,
		Volume:                    historyData.Volume,
		Turnover:                  historyData.Turnover,
		TotalMarketCapitalization: historyData.TotalMarketCapitalization,
		MarketCapitalization:      historyData.MarketCapitalization,
	}
}

func ConvertToResults(datas []model.StockHistoryData) []StockHistoryDataResult {
	results := make([]StockHistoryDataResult, len(datas))

	for index, value := range datas {
		results[index] = ConvertToResult(value)
	}
	return results
}

func ConvertStockResult(stock model.Stock) StockResult {
	return StockResult{
		StockId:   stock.StockId,
		StockCode: stock.StockCode,
		StockName: stock.StockName,
		Type:      stock.Type,
		StartDate: stock.StartDate,
	}
}

func ConvertStockResults(stocks []model.Stock) []StockResult {
	results := make([]StockResult, len(stocks))
	for index, value := range stocks {
		results[index] = ConvertStockResult(value)
	}
	return results
}
