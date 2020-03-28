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
	"terrorsaur/model"
)

//
// @date   2020-02-27 14:47:47
// @author 刘旭
// 切割视频任务的操作

const (
	VideoFileTableName string = "video_file"
	CutVideoTableName  string = "cut_video"
	SourceRootPath     string = "/media/liuxu/LiuXu/crow/source"
	TargetRootPath     string = "/media/liuxu/LiuXu/crow/result"
	FramePath          string = "/poster"
	VideoPath          string = "/video"
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
	c.Db.Table(CutVideoTableName).Where("status!=?", 3).Find(&cutVideos)
	return cutVideos
}

func (c *CutVideoTask) startTask(cutVideos []model.CutVideo) {
	for _, value := range cutVideos {
		c.ExecuteTask(value)
	}
}

func (c *CutVideoTask) ExecuteTask(video model.CutVideo) {
	var videoFile model.VideoFile
	c.Db.Table(VideoFileTableName).Where("video_id=?", video.ParentId).First(&videoFile)

	targetPath := path.Join(TargetRootPath, VideoPath, video.Name)

	commandFormat := "ffmpeg -ss %s -to %s -i %s -c:v libx264 -c:a aac -strict experimental -b:a 98k %s"
	command := fmt.Sprintf(commandFormat, video.StartTime, video.EndTime, videoFile.FilePath, targetPath)
	log.Printf("切割视频执行命令:%s\n", command)
	err := execCommand(command)
	if err != nil {
		log.Printf("截取视频失败!，视频信息是:%s\n", targetPath)
	} else {
		log.Printf("截取视频成功!,视频是:%s\n", targetPath)
		c.Db.Table(CutVideoTableName).Where("cut_id=?", video.CutId).Update(model.CutVideo{FilePath: targetPath, Status: 3})
	}
}

type VideoFileTask struct {
	Db *gorm.DB
}

func (v *VideoFileTask) Start() {
	videoFiles := v.walkFiles(SourceRootPath)
	v.registerVideoFiles(videoFiles)
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

func (v *VideoFileTask) registerVideoFiles(videoFiles []model.VideoFile) []model.VideoFile {
	var count int
	for _, value := range videoFiles {
		v.Db.Table(VideoFileTableName).Where("file_path=?", value.FilePath).Count(&count)
		if count == 0 {
			v.Db.Table(VideoFileTableName).Create(&value)
			go v.takeFrame(value)
		}
	}
	return videoFiles
}

func (v *VideoFileTask) takeFrame(value model.VideoFile) {
	commandFormat := "ffmpeg -ss 3 -i %s -vf \"select=gt(scene\\,0.4)\" -frames:v 1 -vsync vfr -vf fps=fps=1/600 %s/%d.jpg"
	frameDir := filepath.Join(SourceRootPath, FramePath)
	err := os.MkdirAll(frameDir, os.ModePerm)
	if err != nil {
		log.Printf("err")
	}
	command := fmt.Sprintf(commandFormat, value.FilePath, frameDir, value.VideoId)
	log.Printf("执行takeFrame命令:%s", command)
	_ = execCommand(command)
}

func execCommand(command string) error {
	cmd := exec.Command("/bin/bash", "-c", command)
	err := cmd.Run()
	if err != nil {
		log.Print(err)
	}
	return err
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
	/*var videoFileTask = VideoFileTask{
		Db: libs.Db,
	}
	videoFileTask.Start()*/
}
