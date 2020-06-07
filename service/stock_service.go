package service

import (
	"github.com/jinzhu/gorm"
	"terrorsaur/model"
	"terrorsaur/result"
)

//
// @date   2020-04-13 11:06:51
// @author 刘旭
//
// 有关于股票的服务接口
//

const (
	StockHistoryDataTableName string = "stock_history_data"
	StockTableName            string = "stock"
)

type StockService struct {
	Db *gorm.DB
}

func (s *StockService) LoadHistory(stockCode string) []result.StockHistoryDataResult {
	datas := make([]model.StockHistoryData, 0)
	s.Db.Table(StockHistoryDataTableName).Order("date_time", true).Where("stock_code=?", stockCode).Find(&datas)
	return result.ConvertToResults(datas)
}

func (s *StockService) FindStock(condition string) []result.StockResult {
	data := make([]model.Stock, 0)
	likeCondition := "%" + condition + "%"
	s.Db.Table(StockTableName).Where("stock_name like ? or stock_code like ?", likeCondition, likeCondition).Find(&data)
	return result.ConvertStockResults(data)
}
