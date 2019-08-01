package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/core/router"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"log"
	"strconv"
	"terrorsaur/libs"
	"terrorsaur/middleware"
	"terrorsaur/rest"
)

func main() {
	fmt.Println("Terrorsaur Server is starting...")

	app := route()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())

	err := app.Run(iris.Addr(libs.Conf.Server.Domain+":"+strconv.Itoa(libs.Conf.Server.Port)), iris.WithoutServerError(iris.ErrServerClosed))
	if err != nil {
		log.Fatalf("Server start error failed %s", err)
	}
}

func route() (app *iris.Application) {
	app = iris.New()

	v1 := app.Party("/v1", middleware.LogRequestInformationHandler).AllowMethods(iris.MethodOptions)
	{
		v1.PartyFunc("/api/rattrap/ancient-article", func(articleParty router.Party) {
			articleParty.PartyFunc("/user", func(userParty router.Party) {
			})
			articleParty.PartyFunc("/article-type", func(articleTypeParty router.Party) {
				articleTypeParty.Get("/types", rest.ArticleTypePageData)
				articleTypeParty.Get("/{typeId:int}/information", rest.ArticleTypeInformation)
			})
			articleParty.PartyFunc("/article", func(articleParty router.Party) {
			})
		})
	}
	return app
}

func init() {
	fmt.Println("初始化一些配置和数据库信息的操作")
	libs.Db = libs.InitDB()
}
