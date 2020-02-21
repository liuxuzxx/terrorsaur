package service

import (
	"os"
	"path/filepath"
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
	_ = filepath.Walk(rootPath, func(path string, fileInfo os.FileInfo, err error) error {

		videoFileList = append(videoFileList, model.VideoFile{
			FilePath:  path,
			VideoType: 0,
			FileName:  fileInfo.Name(),
			Size:      fileInfo.Size(),
		})
		return nil
	})
	return videoFileList
}
