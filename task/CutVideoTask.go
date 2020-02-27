package task

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"path"
	"terrorsaur/libs"
	"terrorsaur/model"
	"terrorsaur/service"
)

//
// @date   2020-02-27 14:47:47
// @author 刘旭
// 切割视频任务的操作

const (
	VideoFile string = "video_file"
	CutVideo  string = "cut_video"
)

type CutVideoTask struct {
	Db *gorm.DB
}

func (c *CutVideoTask) Start() {
	tasks := c.loadTask()
	c.startTask(tasks)
}

func (c *CutVideoTask) loadTask() []model.CutVideo {
	var cutVideos []model.CutVideo
	c.Db.Table(CutVideo).Where("status!=?", 3).Find(&cutVideos)
	return cutVideos
}

func (c *CutVideoTask) startTask(cutVideos []model.CutVideo) {
	for _, value := range cutVideos {
		c.executeTask(value)
	}
}

func (c *CutVideoTask) executeTask(video model.CutVideo) {
	var videoFile model.VideoFile
	c.Db.Table(VideoFile).Where("video_id=?", video.ParentId).First(&videoFile)

	targetPath := path.Join(service.TargetRootPath, video.Name)

	commandFormat := "ffmpeg -ss %s -t %s -i %s -c:v libx264 -c:a aac -strict experimental -b:a 98k %s"
	command := fmt.Sprintf(commandFormat, video.StartTime, video.EndTime, videoFile.FilePath, targetPath)
	log.Printf("执行的命令字符串是:%s\n", command)
}

func init() {
	log.Printf("启动视频剪切的任务")
	var cutVideoTask = CutVideoTask{
		Db: libs.Db,
	}
	cutVideoTask.Start()
}
