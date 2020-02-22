package libs

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"strconv"
)

var Db *gorm.DB

func InitDB() *gorm.DB {
	//user:password@tcp(host:port)/dbname?charset=utf8&parseTime=True&loc=Local
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", Conf.DataSource.UserName, Conf.DataSource.Password, Conf.DataSource.Host, strconv.Itoa(Conf.DataSource.Port), Conf.DataSource.DatabaseName, Conf.DataSource.Charset)
	log.Printf("查看下具体的连接字符串:%s\n", connString)
	db, err := gorm.Open("mysql", connString)
	if err != nil {
		log.Panic(err)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(Conf.DataSource.MaxIdleConnections)
	db.DB().SetMaxOpenConns(Conf.DataSource.MaxOpenConnections)
	return db
}

func init() {
	log.Printf("初始化数据库的连接信息\n")
	Db = InitDB()
}
