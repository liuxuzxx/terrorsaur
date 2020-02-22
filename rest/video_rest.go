package rest

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/nareix/joy4/format"
	"io/ioutil"
	"terrorsaur/common"
	"terrorsaur/service"
)

//
// @Date 2020-02-21 09:36-30
// @Author 刘旭
//go开发视频服务器的实验
func init() {
	format.RegisterAll()
}

func VideoPlayer(ctx iris.Context) {
	videoId, _ := ctx.Params().GetInt64("videoId")
	videoFileResult := service.FetchVideoFile(videoId)
	readFile, err := ioutil.ReadFile(videoFileResult.FilePath)
	if err != nil {
		fmt.Println("读取信息出现错误")
	}
	ctx.ContentType("video/mp4")
	ctx.Write(readFile)
}

func VideoFiles(ctx iris.Context) {
	files := service.FetchAllVideoFiles()
	_, _ = ctx.JSON(common.Success(files))
}
