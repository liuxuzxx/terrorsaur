package task

//
// @date   2020-03-28 17:12:05
// @author 刘旭
//
// 股票的任务执行
//

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"terrorsaur/model"
	"terrorsaur/request"
)

const (
	StockTableName = "stock"
)

type StockTask struct {
	StockCount int64
	Db         *gorm.DB
}

func (s *StockTask) crawlStockHtml(url string) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("请求:%s地址失败", url)
	}
	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatalf("解析文档错误")
	}
	content := document.Text()
	content = strings.ReplaceAll(content, "jQuery112404885528144188753_1585461784221(", "")
	content = strings.ReplaceAll(content, ");", "")
	var stockJson request.StockJson
	err = json.Unmarshal([]byte(content), &stockJson)

	if err != nil {
		log.Printf("解析json字符串出现错误:%s", content)
	}
	for _, stock := range stockJson.Data.Diff {
		s.persistenceStock(request.Stock{
			Code: stock.F12,
			Name: stock.F14,
		})
		s.StockCount++
	}
}

func (s *StockTask) persistenceStock(stockRequest request.Stock) {
	stock := model.Stock{
		StockCode: stockRequest.Code,
		StockName: stockRequest.Name,
	}
	if strings.HasPrefix(stockRequest.Code, "6") {
		stock.Type = "沪"
	} else {
		stock.Type = "深"
	}
	var count int
	s.Db.Table(StockTableName).Where("stock_code=?", stock.StockCode).Count(&count)
	if count == 0 {
		s.Db.Table(StockTableName).Create(&stock)
	} else {
		log.Printf("股票信息:%s %s 已经插入到数据库当中了!", stock.StockName, stock.StockCode)
	}
}

func (s *StockTask) execute() {
	var urlTemplate = "http://29.push2.eastmoney.com/api/qt/clist/get?cb=jQuery112404885528144188753_1585461784221&pn=%d&pz=20&po=1&np=1&ut=bd1d9ddb04089700cf9c27f6f7426281&fltt=2&invt=2&fid=f3&fs=m:0+t:6,m:0+t:13,m:0+t:80,m:1+t:2,m:1+t:23&fields=f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21,f23,f24,f25,f22,f11,f62,f128,f136,f115,f152&_=1585461784239"
	s.doStockExecute(urlTemplate, 197)
	s.crawStockHistoryData()
}

func (s *StockTask) doStockExecute(urlTemplate string, pageCount int) {
	/*for pageNumber := 1; pageNumber <= pageCount; pageNumber++ {
		url := fmt.Sprintf(urlTemplate, pageNumber)
		s.crawlStockHtml(url)
	}
	log.Printf("爬取的股票数量是:%d\n", s.StockCount)*/
}

func (s *StockTask) crawStockHistoryData() {
	historyDataUrlTemplate := "http://quotes.money.163.com/service/chddata.html?code=%s&start=%s&end=20200401&fields=TCLOSE;HIGH;LOW;TOPEN;LCLOSE;CHG;PCHG;TURNOVER;VOTURNOVER;VATURNOVER;TCAP;MCAP"
	stocks := s.fetchAllStocks()
	for _, stock := range stocks {
		code := stock.StockCode
		if len(stock.StartDate) < 10 {
			continue
		}
		if strings.HasPrefix(code, "6") {
			code = "0" + code
		} else {
			code = "1" + code
		}
		url := fmt.Sprintf(historyDataUrlTemplate, code, strings.ReplaceAll(stock.StartDate, "-", ""))
		s.downloadFile(url, filepath.Join("/home/liuxu/Documents/stock", stock.StockCode+".csv"))
	}
}

func (s *StockTask) downloadFile(url string, filePath string) {
	command := "wget \"" + url + "\" -O " + filePath + "\n"
	fmt.Printf("url是:%s\n", command)
	s.appendFile(command)
}

func (s *StockTask) appendFile(content string) {
	file := "/home/liuxu/Documents/stock-history.sh"
	f, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	if err != nil {
		fmt.Printf("Cannot open file %s!\n", file)
		return
	}
	defer f.Close()
	_, _ = f.WriteString(content)
}

func (s *StockTask) fetchAllStocks() []model.Stock {
	var stocks []model.Stock
	s.Db.Table(StockTableName).Find(&stocks)
	return stocks
}

func init() {
	log.Printf("开始爬取股票数据.......")
	/*var task = StockTask{
		StockCount: int64(0),
		Db:         libs.Db,
	}
	task.execute()*/
}
