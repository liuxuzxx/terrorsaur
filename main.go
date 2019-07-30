package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/kataras/iris"
	"github.com/kataras/iris/core/router"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/spf13/viper"
	"log"
	"strconv"
	"terrorsaur/rest"
)

func main() {
	fmt.Println("Terrorsaur Server is starting...")

	app := route()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())

	err := app.Run(iris.Addr(config.Server.Domain+":"+strconv.Itoa(config.Server.Port)), iris.WithoutServerError(iris.ErrServerClosed))
	if err != nil {
		log.Fatalf("Server start error failed %s", err)
	}
}

func route() (app *iris.Application) {
	app = iris.New()

	v1 := app.Party("/v1").AllowMethods(iris.MethodOptions)
	{
		v1.PartyFunc("/api/rattrap/ancient-article", func(articleParty router.Party) {
			articleParty.PartyFunc("/user", func(userParty router.Party) {
			})
			articleParty.PartyFunc("/article-type", func(articleTypeParty router.Party) {
				articleParty.Get("/types", rest.ArticleTypePageData)
				articleParty.Get("/{typeId:int}", rest.ArticleTypeInformation)
			})
			articleParty.PartyFunc("/article", func(articleParty router.Party) {
			})
		})
	}
	return app
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
