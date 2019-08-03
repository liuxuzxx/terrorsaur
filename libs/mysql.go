package libs

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"strconv"
)

var Db *gorm.DB

func InitDB() *gorm.DB {
	//user:password@tcp(host:port)/dbname?charset=utf8&parseTime=True&loc=Local
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", Conf.DataSource.UserName, Conf.DataSource.Password, Conf.DataSource.Host, strconv.Itoa(Conf.DataSource.Port), Conf.DataSource.DatabaseName, Conf.DataSource.Charset)
	fmt.Println("查看下具体的连接字符串:", connString)
	db, err := gorm.Open("mysql", connString)
	if err != nil {
		log.Panic(err)
	}
	db.SingularTable(true)                                      //全局设置表名不可以为复数形式。
	db.DB().SetMaxIdleConns(Conf.DataSource.MaxIdleConnections) //空闲时最大的连接数
	db.DB().SetMaxOpenConns(Conf.DataSource.MaxOpenConnections) //最大的连接数
	return db
}
