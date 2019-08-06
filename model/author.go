package model

type Author struct {
	AuthorId int    `gorm:"column:author_id"`
	Name     string `gorm:"column:name"`
	Dynasty  string `gorm:"column:dynasty"`
	Detail   string `gorm:"column:detail"`
}
