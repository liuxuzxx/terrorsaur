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
	"log"
	"net/http"
	"strings"
)

type StockTask struct {
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
	content = strings.ReplaceAll(content, "jQuery1124008598677345336991_1585381151677(", "")
	content = strings.ReplaceAll(content, ");", "")
	var jsonObject map[string]interface{}
	err = json.Unmarshal([]byte(content), &jsonObject)
	diff := jsonObject["data"].(map[string]interface{})["diff"]
	diff = diff.([]interface{})

	log.Printf("diff:%d", diff)
	if err != nil {
		log.Printf("解析json字符串出现错误")
	}
}

func (s *StockTask) execute() {
	var urlTemplate = "http://88.push2.eastmoney.com/api/qt/clist/get?cb=jQuery1124008598677345336991_1585381151677&pn=%d&pz=20&po=1&np=1&ut=bd1d9ddb04089700cf9c27f6f7426281&fltt=2&invt=2&fid=f3&fs=m:0+t:6,m:0+t:13,m:0+t:80&fields=f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21,f23,f24,f25,f22,f11,f62,f128,f136,f115,f152&_=1585381151793"
	s.doExecute(urlTemplate, 114)
}

func (s *StockTask) doExecute(urlTemplate string, pageCount int) {
	for pageNumber := 1; pageNumber <= pageCount; pageNumber++ {
		url := fmt.Sprintf(urlTemplate, pageNumber)
		s.crawlStockHtml(url)
	}
}

func init() {
	log.Printf("开始爬取股票数据.......")
	var task = StockTask{}
	task.execute()
}
