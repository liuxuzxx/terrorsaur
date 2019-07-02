package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/core/router"
	"github.com/kataras/iris/middleware/logger"
	recover2 "github.com/kataras/iris/middleware/recover"
	"mime/multipart"
	"strings"
)

const maxSize = 5 << 20

func main() {
	fmt.Println("直接开启Iris的Go的web编程")
	app := iris.New()
	app.Use(recover2.New())
	app.Use(logger.New())

	app.PartyFunc("/api/system", func(system router.Party) {
		system.Get("/information", systemInformationHandler)
		system.Get("/name", systemNameHandler)
	})

	app.PartyFunc("/api/user", func(users router.Party) {
		users.Get("/information", userInformationHandler)
		users.Get("/{userId:int}/status", userStatusHandler)
		users.Get("/{userId:int}/birthday", userBirthdayHandler)
	})

	app.PartyFunc("/api/video", func(video router.Party) {
		video.Post("/upload", iris.LimitRequestBodySize(maxSize), uploadVideoHandler)
	})

	app.Run(iris.Addr(":12309"), iris.WithConfiguration(iris.Configuration{
		Charset: "UTF-8",
	}))
}

func systemInformationHandler(ctx iris.Context) {
	ctx.JSON(iris.Map{"systemName": "Go-Iris系统", "status": "OK"})
}

func systemNameHandler(ctx iris.Context) {
	ctx.JSON(iris.Map{"systemName": "Go-Iris系统"})
}

func userInformationHandler(ctx iris.Context) {
	ctx.JSON(iris.Map{"userName": "rootAdmin", "age": 20, "birthday": "1990-10-07"})
}

func userStatusHandler(ctx iris.Context) {
	userId := ctx.Params().Get("userId")
	ctx.JSON(iris.Map{"status": "OK", "userId": userId})
}

func userBirthdayHandler(ctx iris.Context) {
	userId := ctx.Params().Get("userId")
	ctx.JSON(iris.Map{"userId": userId, "birthday": "1990-10-06"})
}

func uploadVideoHandler(ctx iris.Context) {
	ctx.UploadFormFiles("F:/", beforeSave)
}

func beforeSave(ctx iris.Context, file *multipart.FileHeader) {
	ip := ctx.RemoteAddr()
	ip = strings.Replace(ip, ".", "_", -1)
	ip = strings.Replace(ip, ":", "_", -1)

	file.Filename = ip + "-" + file.Filename
}
