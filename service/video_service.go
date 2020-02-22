package service

import (
	"os"
	"path"
	"path/filepath"
	"strings"
	"terrorsaur/model"
)

//
// @Date 2020-02-21 15:01-56
// @Author 刘旭
// 视频服务操作

func VideoFiles(rootPath string) []model.VideoFile {
	return walkFiles(rootPath)
}

func walkFiles(rootPath string) []model.VideoFile {
	var videoFileList []model.VideoFile
	_ = filepath.Walk(rootPath, func(filePath string, fileInfo os.FileInfo, err error) error {
		videoType, ok := verifyVideoType(fileInfo.Name())
		if ok {
			videoFileList = append(videoFileList, model.VideoFile{
				FilePath:  filePath,
				VideoType: videoType,
				FileName:  fileInfo.Name(),
				Size:      fileInfo.Size(),
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
	".JAVA": {
		TypeCode: 5,
		TypeName: "Java文件类型",
	},
	"UNKNOWN": {
		TypeCode: 4,
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
