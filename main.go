package main

import (
	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"log"
	"strconv"
	_ "terrorsaur/docs"
	"terrorsaur/libs"
	"terrorsaur/middleware"
	"terrorsaur/rest"
)

// @title Terrorsaur-鸵鸟勇士
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 172.16.21.207:12309
// @BasePath /v1
func main() {
	log.Printf("Terrorsaur Server is starting...")

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
	config := &swagger.Config{
		URL: "http://" + libs.Conf.Server.Domain + ":" + strconv.Itoa(libs.Conf.Server.Port) + "/swagger/doc.json", //The url pointing to API definition
	}
	app.Get("/swagger/{any:path}", swagger.CustomWrapHandler(config, swaggerFiles.Handler))

	v1 := app.Party("/v1",
		middleware.LogRequestInformationHandler).AllowMethods(iris.MethodOptions)
	{
		v1.PartyFunc("/api/rattrap/ancient-article", func(articleParty router.Party) {
			articleParty.PartyFunc("/plate", func(plateParty router.Party) {
				plateParty.Get("", rest.AncientPlateInformation)
			})
			articleParty.PartyFunc("/user", func(userParty router.Party) {
			})
			articleParty.PartyFunc("/article-type", func(articleTypeParty router.Party) {
				articleTypeParty.Get("/types", rest.ArticleTypePageData)
				articleTypeParty.Get("/{typeId:int}/information", rest.ArticleTypeInformation)
			})
			articleParty.PartyFunc("/article", func(articleParty router.Party) {
				articleParty.Get("/{typeId:int}/articles", rest.ArticlesInformation)
				articleParty.Get("/{articleId:int}/detail", rest.ArticleDetailInformation)
			})
			articleParty.PartyFunc("/author", func(authorParty router.Party) {
				authorParty.Get("/{authorId:int}/information", rest.AuthorInformation)
			})
			articleParty.PartyFunc("/phrase", func(phraseParty router.Party) {
				phraseParty.Get("/information", rest.IdiomsInformation)
			})
		})
		v1.PartyFunc("/api/rattrap/dictionary", func(dictionaryParty router.Party) {
			dictionaryParty.PartyFunc("/dictionary-types", func(dictionaryTypeParty router.Party) {
				dictionaryTypeParty.Get("", rest.DictionaryTypeInformation)
			})
		})
	}
	return app
}

func init() {
	log.Println("初始化一些配置和数据库信息的操作")
	libs.Db = libs.InitDB()
}
