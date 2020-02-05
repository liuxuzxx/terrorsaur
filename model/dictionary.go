package model

//
// @Date 2019-12-09 18:10-03
// @Author 刘旭
//主要是应用于字典方面的
//

//字典类型信息对象
type DictionaryType struct {
	Id             int    `gorm:"column:id"`
	DictionaryCode string `gorm:"column:dictionary_code"`
	DictionaryName string `gorm:"column:dictionary_name"`
}

//单个字体信息实体类
type Dictionary struct {
	Id               int    `gorm:"column:id"`
	ChineseCharacter string `gorm:"column:chinese_character"`
	Explanation      string `gorm:"column:explanation"`
	Source           string `gorm:"column:source"`
	DictionaryType   int    `gorm:"column:dictionary_type"`
}
