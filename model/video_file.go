package model

//
// @Date 2020-02-21 14:54-32
// @Author 刘旭
//主要是存放有关于视频的实体类的对象

type VideoFile struct {
	VideoId  int64  `gorm:"column:video_id"`
	FilePath string `gorm:"column:file_path"`
	FileName string `gorm:"column:file_name"`
	Size     int64  `gorm:"column:size"`
}

type VideoType struct {
	TypeCode int8
	TypeName string
}
