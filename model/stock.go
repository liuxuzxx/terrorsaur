package model

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
}
