package model

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ArticleType struct {
	TypeId    int    `gorm:"column:type_id"`
	TypeName  string `gorm:"column:type_name"`
	Detail    string `gorm:"column:detail"`
	TypeOrder int    `gorm:"column:type_order"`
}
