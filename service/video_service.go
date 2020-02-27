package service

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"terrorsaur/libs"
	"terrorsaur/model"
	"terrorsaur/result"
)

//
// @Date 2020-02-21 15:01-56
// @Author 刘旭
// 视频服务操作

const (
	VideoFileTableName string = "video_file"
	RootPath           string = "/media/liuxu/LiuXu/crow/source"
	TargetRootPath     string = "/media/liuxu/LiuXu/crow/target"
	CutVideo           string = "cut_video"
)

func walkFiles(rootPath string) []model.VideoFile {
	var videoFileList []model.VideoFile
	_ = filepath.Walk(rootPath, func(filePath string, fileInfo os.FileInfo, err error) error {
		_, ok := verifyVideoType(fileInfo.Name())
		if ok {
			videoFileList = append(videoFileList, model.VideoFile{
				FilePath: filePath,
				FileName: fileInfo.Name(),
				Size:     fileInfo.Size(),
			})
		}
		return nil
	})
	return videoFileList
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

func FetchAllVideoFiles() []result.VideoFileResult {
	var videoFiles []model.VideoFile
	libs.Db.Table(VideoFileTableName).Find(&videoFiles)
	return result.ConvertVideoFileToResults(videoFiles)
}

func FetchVideoFile(videoId int64) result.VideoFileResult {
	var videoFile model.VideoFile
	libs.Db.Table(VideoFileTableName).Where("video_id=?", videoId).First(&videoFile)
	return result.ConvertVideoFileToResult(videoFile)
}

func RegisterCutVideo(request result.CutVideoRequest) {
	cutVideo := result.ConvertCutVideoRequestTo(request)
	var videoFile model.VideoFile
	libs.Db.Table(VideoFileTableName).Where("video_id=?", request.ParentId).First(&videoFile)
	cutVideo.Name = strings.Join([]string{request.StartTime, request.EndTime, videoFile.FileName}, "-")
	libs.Db.Table(CutVideo).Create(&cutVideo)
}

func FetchAllCutById(parentId int64) []result.CutVideoResult {
	var videos []model.CutVideo
	libs.Db.Table(CutVideo).Where("parent_id=?", parentId).Find(&videos)
	return result.ConvertCutVideoToResults(videos)
}

func registerVideoFiles(videoFiles []model.VideoFile) {
	for _, value := range videoFiles {
		libs.Db.Table(VideoFileTableName).Create(&value)
	}
}

func init() {
	log.Printf("初始化视频服务数据!")
	//videoFiles := walkFiles(RootPath)
	//registerVideoFiles(videoFiles)
}
