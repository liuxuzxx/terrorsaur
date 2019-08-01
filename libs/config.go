package libs

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
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
	Ip           string
	Port         int
	UserName     string
	Password     string
	DatabaseName string
}

var Conf Config

func init(){
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Update the config file")
	})
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Print("Parse the config file is error!")
	}
	err = viper.Unmarshal(&Conf)
}
