package model

import "time"

//
// @date   2020-03-29 22:31:25
// @author 刘旭
//
// 主要是放置股票的数据库对象
//

type Stock struct {
	StockId   int64  `gorm:"column:stock_id"`
	StockCode string `gorm:"column:stock_code"`
	StockName string `gorm:"column:stock_name"`
	Type      string `gorm:"column:type"`
	StartDate string `gorm:"column:start_date"`
}

type StockHistoryData struct {
	HistoryId                 int64     `gorm:"column:history_id"`
	StockCode                 string    `gorm:"column:stock_code"`
	DateTime                  time.Time `gorm:"column:date_time"`
	StartPrice                float64   `gorm:"column:start_price"`
	EndPrice                  float64   `gorm:"column:end_price"`
	YesterdayEndPrice         float64   `gorm:"column:yesterday_end_price"`
	MaxPrice                  float64   `gorm:"column:max_price"`
	MinPrice                  float64   `gorm:"column:min_price"`
	RiseFall                  float64   `gorm:"column:rise_fall"`
	TurnoverRate              float64   `gorm:"column:turnover_rate"`
	Volume                    int64     `gorm:"column:volume"`
	Turnover                  float64   `gorm:"column:turnover"`
	TotalMarketCapitalization float64   `gorm:"column:total_market_capitalization"`
	MarketCapitalization      float64   `gorm:"column:market_capitalization"`
}
