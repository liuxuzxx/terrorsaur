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

type CutVideoRequest struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	ParentId  int64  `json:"parentId"`
}

type CutVideoResult struct {
	CutId     int64  `json:"cutId"`
	ParentId  int64  `json:"parentId"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	Name      string `json:"name"`
	Status    int64  `json:"status"`
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

func ConvertCutVideoRequestTo(request CutVideoRequest) model.CutVideo {
	return model.CutVideo{
		ParentId:  request.ParentId,
		StartTime: request.StartTime,
		EndTime:   request.EndTime,
		Status:    1,
	}
}

func ConvertCutVideoToResult(video model.CutVideo) CutVideoResult {
	return CutVideoResult{
		CutId:     video.CutId,
		ParentId:  video.ParentId,
		StartTime: video.StartTime,
		EndTime:   video.EndTime,
		Name:      video.Name,
		Status:    video.Status,
	}
}

func ConvertCutVideoToResults(videos []model.CutVideo) []CutVideoResult {
	var results = make([]CutVideoResult, len(videos))

	for index, value := range videos {
		results[index] = ConvertCutVideoToResult(value)
	}
	return results
}
