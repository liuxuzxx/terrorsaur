package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/spf13/viper"
	"log"
	"strconv"
)

func main() {
	fmt.Println("Terrorsaur Server is starting...")

	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())

	err := app.Run(iris.Addr(config.Server.Domain+":"+strconv.Itoa(config.Server.Port)), iris.WithoutServerError(iris.ErrServerClosed))
	if err != nil {
		log.Fatalf("Server start error failed %s", err)
	}
}

var config Config

func init() {
	fmt.Println("Start init the web config information!")
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
	err = viper.Unmarshal(&config)
}

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
