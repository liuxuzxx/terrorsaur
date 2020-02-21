package model

//
// @Date 2020-02-21 14:54-32
// @Author 刘旭
//主要是存放有关于视频的实体类的对象

type VideoFile struct {
	FilePath  string
	VideoType VideoType
	FileName  string
	Size      int64
}

type VideoType struct {
	TypeCode int8
	TypeName string
}
