package task

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"terrorsaur/libs"
	"terrorsaur/model"
	"terrorsaur/service"
)

//
// @date   2020-02-27 14:47:47
// @author 刘旭
// 切割视频任务的操作

type CutVideoTask struct {
	Db *gorm.DB
}

func (c *CutVideoTask) Start() {
	tasks := c.loadTask()
	c.startTask(tasks)
}

func (c *CutVideoTask) loadTask() []model.CutVideo {
	var cutVideos []model.CutVideo
	c.Db.Table(service.CutVideoTableName).Where("status!=?", 3).Find(&cutVideos)
	return cutVideos
}

func (c *CutVideoTask) startTask(cutVideos []model.CutVideo) {
	for _, value := range cutVideos {
		c.executeTask(value)
	}
}

func (c *CutVideoTask) executeTask(video model.CutVideo) {
	var videoFile model.VideoFile
	c.Db.Table(service.VideoFileTableName).Where("video_id=?", video.ParentId).First(&videoFile)

	targetPath := path.Join(service.TargetRootPath, video.Name)

	commandFormat := "ffmpeg -ss %s -to %s -i %s -c:v libx264 -c:a aac -strict experimental -b:a 98k %s"
	command := fmt.Sprintf(commandFormat, video.StartTime, video.EndTime, videoFile.FilePath, targetPath)
	log.Printf("执行的命令字符串是:%s\n", command)
	//execCommand(command)
}

type VideoFileTask struct {
	Db *gorm.DB
}

func (v *VideoFileTask) Start() {
	videoFiles := v.walkFiles(service.SourceRootPath)
	//v.registerVideoFiles(videoFiles)
	v.takeFrame(videoFiles)
}

func (v *VideoFileTask) walkFiles(rootPath string) []model.VideoFile {
	var videoFileList []model.VideoFile
	_ = filepath.Walk(rootPath, func(filePath string, fileInfo os.FileInfo, err error) error {
		if !fileInfo.IsDir() {
			_, ok := verifyVideoType(fileInfo.Name())
			if ok {
				videoFileList = append(videoFileList, model.VideoFile{
					FilePath: filePath,
					FileName: fileInfo.Name(),
					Size:     fileInfo.Size(),
				})
			}
		}
		return nil
	})
	return videoFileList
}

func (v *VideoFileTask) registerVideoFiles(videoFiles []model.VideoFile) {
	for _, value := range videoFiles {
		v.Db.Table(service.VideoFileTableName).Create(&value)
	}
}

func (v *VideoFileTask) takeFrame(videoFiles []model.VideoFile) {
	commandFormat := "ffmpeg -i %s -r 1 -q:v 2 -f image2 %s/%s.jpg"
	for _, value := range videoFiles {
		dir := filepath.Dir(value.FilePath)
		frameDir := filepath.Join(dir, service.FramePath, value.FileName)
		err := os.MkdirAll(frameDir, os.ModePerm)
		if err != nil {
			log.Printf("err")
		}
		command := fmt.Sprintf(commandFormat, value.FilePath, frameDir, "%d")
		log.Printf("执行takeFrame命令:%s", command)
		//execCommand(command)
	}
}

func execCommand(command string) {
	cmd := exec.Command("/bin/bash", "-c", command)
	err := cmd.Run()
	if err != nil {
		log.Print(err)
	}
}

var fileType = map[string]model.VideoType{
	".MP4": {
		TypeCode: 1,
		TypeName: "mp4",
	},
	".FLV": {
		TypeCode: 2,
		TypeName: "FLV",
	},
	".MKV": {
		TypeCode: 3,
		TypeName: "mkv",
	},
	".F4V": {
		TypeCode: 5,
		TypeName: "f4v",
	},
	"AVI": {
		TypeCode: 6,
		TypeName: "avi",
	},
	"WMV": {
		TypeCode: 7,
		TypeName: "wmv",
	},
	"MOV": {
		TypeCode: 8,
		TypeName: "mov",
	},
	"RMVB": {
		TypeCode: 9,
		TypeName: "rmvb",
	},
	"UNKNOWN": {
		TypeCode: -1,
		TypeName: "未知类型",
	},
}

func verifyVideoType(filePath string) (model.VideoType, bool) {
	fileSuffix := path.Ext(filePath)
	videoType, ok := fileType[strings.ToUpper(fileSuffix)]
	if ok {
		return videoType, true
	}
	return fileType["UNKNOWN"], false
}

func init() {
	log.Printf("启动视频剪切的任务")
	var cutVideoTask = CutVideoTask{
		Db: libs.Db,
	}
	cutVideoTask.Start()

	var videoFileTask = VideoFileTask{
		Db: libs.Db,
	}
	videoFileTask.Start()
}
