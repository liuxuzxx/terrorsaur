package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ArticleType struct{
	gorm.Model
	TypeId int
	TypeName string
	Detail string
	TypeOrder int
}


