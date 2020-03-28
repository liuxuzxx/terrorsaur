package model

//
// @Date 2020-02-21 14:54-32
// @Author 刘旭
//主要是存放有关于视频的实体类的对象

type VideoFile struct {
	VideoId  int64  `gorm:"primary_key;column:video_id"`
	FilePath string `gorm:"column:file_path"`
	FileName string `gorm:"column:file_name"`
	Size     int64  `gorm:"column:size"`
	Poster   string `gorm:"column:poster"`
}

type VideoType struct {
	TypeCode int8
	TypeName string
}

type CutVideo struct {
	CutId     int64  `gorm:"column:cut_id"`
	ParentId  int64  `gorm:"column:parent_id"`
	StartTime string `gorm:"column:start_time"`
	EndTime   string `gorm:"column:end_time"`
	Name      string `gorm:"column:name"`
	Status    int64  `gorm:"column:status"`
	Poster    string `gorm:"column:poster"`
	FilePath  string `gorm:"column:file_path"`
}
