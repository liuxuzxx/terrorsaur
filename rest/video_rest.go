package rest

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/nareix/joy4/av"
	"github.com/nareix/joy4/av/avutil"
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
	videoPath := "/home/liuxu/Downloads/mociro-unit.mp4"
	file, _ := avutil.Open(videoPath)
	streams, _ := file.Streams()
	for _, stream := range streams {
		if stream.Type().IsAudio() {
			astream := stream.(av.AudioCodecData)
			fmt.Println(astream.Type(), astream.SampleRate(), astream.SampleFormat(), astream.ChannelLayout())
		} else if stream.Type().IsVideo() {
			vstream := stream.(av.VideoCodecData)
			fmt.Println("是一个video")
			fmt.Println(vstream.Type(), vstream.Width(), vstream.Height())
		}
	}

	file.Close()

	readFile, err := ioutil.ReadFile(videoPath)
	if err != nil {
		fmt.Println("读取信息出现错误")
	}
	ctx.ContentType("video/mp4")
	ctx.Write(readFile)
}

func VideoFiles(ctx iris.Context) {
	rootPath := "/media/liuxu/data/leonard/relax"
	files := service.VideoFiles(rootPath)
	_, _ = ctx.JSON(common.Success(files))
}
