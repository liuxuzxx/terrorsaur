package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func main() {
	fmt.Println("Terrorsaur Server启动了......")
	fmt.Println(config.Server.Name)
	fmt.Println(config.Server.Port)
	fmt.Println(config.DataSource)
}

var config Config

func init() {
	fmt.Println("Start init the web config information!")
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("修改了配置文件")
	})
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Print("解析出现错误")
	}
	err = viper.Unmarshal(&config)
}

type Config struct{
	Server Server
	DataSource DataSource
}

type Server struct{
	Port int
	Name string
}

type DataSource struct{
	Ip string
	Port int
	UserName string
	Password string
	DatabaseName string
}
