package service

import (
	"path/filepath"
	"strings"
	"terrorsaur/libs"
	"terrorsaur/model"
	"terrorsaur/result"
	"terrorsaur/task"
)

//
// @Date 2020-02-21 15:01-56
// @Author 刘旭
// 视频服务操作

const (
	VideoFileTableName string = "video_file"
	CutVideoTableName  string = "cut_video"
	SourceRootPath     string = "/media/liuxu/LiuXu/crow/source"
	TargetRootPath     string = "/media/liuxu/LiuXu/crow/result"
	FramePath          string = "/poster"
	VideoPath          string = "video"
)

var cutVideoTask = task.CutVideoTask{Db: libs.Db}

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
	libs.Db.Table(CutVideoTableName).Create(&cutVideo)
	go cutVideoTask.ExecuteTask(cutVideo)
}

func FetchAllCutById(parentId int64) []result.CutVideoResult {
	var videos []model.CutVideo
	libs.Db.Table(CutVideoTableName).Where("parent_id=?", parentId).Find(&videos)
	return result.ConvertCutVideoToResults(videos)
}

func FetchHeadFrame(videoId int64) string {
	var videoFile model.VideoFile
	libs.Db.Table(VideoFileTableName).Where("video_id=?", videoId).First(&videoFile)
	return filepath.Join(SourceRootPath, FramePath, videoFile.Poster)
}
