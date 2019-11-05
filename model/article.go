package model

import (
	"time"
)

type ArticleType struct {
	TypeId    int    `gorm:"column:type_id"`
	TypeName  string `gorm:"column:type_name"`
	Detail    string `gorm:"column:detail"`
	TypeOrder int    `gorm:"column:type_order"`
}

type ArticleAttribute struct {
	ArticleId      int    `gorm:"column:article_id"`
	AttributeCode  int    `gorm:"column:attribute_code"`
	AttributeValue string `gorm:"column:attribute_value"`
}

type Article struct {
	ArticleId  int       `gorm:"column:article_id"`
	Title      string    `gorm:"column:title"`
	Content    string    `gorm:"column:content"`
	Comment    string    `gorm:"column:comment"`
	CreateBy   string    `gorm:"column:create_by"`
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateBy   string    `gorm:"column:update_by"`
	UpdateTime time.Time `gorm:"column:update_time"`
}
