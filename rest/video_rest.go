package rest

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/nareix/joy4/format"
	"io"
	"os"
	"strconv"
	"terrorsaur/common"
	"terrorsaur/result"
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
	var startRange, endRange int64

	_, _ = fmt.Sscanf(ctx.GetHeader("Range"), "bytes=%d-%d", &startRange, &endRange)
	videoId, _ := ctx.Params().GetInt64("videoId")
	videoFileResult := service.FetchVideoFile(videoId)

	videoFile, videoFileErr := os.Open(videoFileResult.FilePath)
	if videoFileErr != nil {
		ctx.NotFound()
		return
	}
	videoFileInfo, _ := videoFile.Stat()
	if endRange == 0 {
		endRange = videoFileInfo.Size() - 1
	}
	ctx.Header("Accept-Ranges", "bytes")
	ctx.Header("Content-Length", strconv.FormatInt(endRange-startRange+1, 10))
	ctx.Header("Content-Range", "bytes "+strconv.FormatInt(startRange, 10)+"-"+strconv.FormatInt(endRange, 10)+"/"+strconv.FormatInt(videoFileInfo.Size(), 10))
	ctx.Header("Content-Disposition", "attachment; filename="+videoFileInfo.Name())

	_, seekErr := videoFile.Seek(startRange, 0)
	if seekErr != nil {
		ctx.NotFound()
	}
	ctx.StatusCode(206)
	_, _ = io.CopyN(ctx, videoFile, endRange-startRange+1)
}

func VideoFiles(ctx iris.Context) {
	files := service.FetchAllVideoFiles()
	_, _ = ctx.JSON(common.Success(files))
}

func CutVideoRegister(ctx iris.Context) {
	var cutVideoRequest result.CutVideoRequest
	_ = ctx.ReadJSON(&cutVideoRequest)
	service.RegisterCutVideo(cutVideoRequest)
	_, _ = ctx.JSON(common.Success("成功注册切割视频"))
}

func FetchCutVideos(ctx iris.Context) {
	parentId, _ := ctx.Params().GetInt64("parentId")
	_, _ = ctx.JSON(common.Success(service.FetchAllCutById(parentId)))
}
