package result

import "terrorsaur/model"

//
// @date   2020-02-22 20:10:50
// @author 刘旭
// 视频结果对象

type VideoFileResult struct {
	VideoId  int64  `json:"videoId"`
	FilePath string `json:"filePath"`
	FileName string `json:"fileName"`
	Size     int64  `json:"size"`
}

func ConvertVideoFileToResult(videoFile model.VideoFile) VideoFileResult {
	return VideoFileResult{
		VideoId:  videoFile.VideoId,
		FilePath: videoFile.FilePath,
		FileName: videoFile.FileName,
		Size:     videoFile.Size,
	}
}

func ConvertVideoFileToResults(videoFiles []model.VideoFile) []VideoFileResult {
	results := make([]VideoFileResult, len(videoFiles))
	for index, value := range videoFiles {
		results[index] = ConvertVideoFileToResult(value)
	}
	return results
}
