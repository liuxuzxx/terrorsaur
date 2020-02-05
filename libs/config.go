package libs

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Server     Server
	DataSource DataSource
}

type Server struct {
	Domain string
	Port   int
	Name   string
}

type DataSource struct {
	Host               string
	Port               int
	UserName           string
	Password           string
	DatabaseName       string
	Charset            string
	MaxIdleConnections int
	MaxOpenConnections int
}

var Conf Config

func init() {
	log.Println("Load config information from config directory!")
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Update the config file")
	})
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Parse the config file is error!")
	}
	err = viper.Unmarshal(&Conf)
}
