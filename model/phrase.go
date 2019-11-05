package model

//
//主要是短语类型，因为诗词都是那种长篇大论的形式，短语往往给人一种醍醐灌顶的感觉
//
type Idioms struct {
	Id             int    `gorm:"column:id"`
	Term           string `gorm:"column:term"`
	Pronunciation  string `gorm:"column:pronunciation"`
	Interpretation string `gorm:"column:interpretation"`
	Source         string `gorm:"column:source"`
	Example        string `gorm:"column:example"`
}
