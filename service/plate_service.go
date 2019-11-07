package service

import (
	"terrorsaur/libs"
	"terrorsaur/model"
	"terrorsaur/result"
)

//
//板块的服务
//
const (
	AncientPlateTableName string = "ancient_plate"
)

func FetchAllPlates() []result.AncientPlateResult {
	var ancientPlates []model.AncientPlate
	libs.Db.Table(AncientPlateTableName).Find(&ancientPlates)
	return result.ConvertAncientPlateResults(ancientPlates)
}
