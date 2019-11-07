package model

//
//古代的板块信息
//
type AncientPlate struct {
	PlateId    int    `gorm:"column:plate_id"`
	PlateName  string `gorm:"column:plate_name"`
	Detail     string `gorm:"column:detail"`
	PlateOrder int    `gorm:"column:plate_order"`
}
