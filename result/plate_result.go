package result

import "terrorsaur/model"

//
//古代板块的结果对象信息
//
type AncientPlateResult struct {
	PlateId    int    `json:"plateId"`
	PlateName  string `json:"plateName"`
	Detail     string `json:"detail"`
	PlateOrder int    `json:"plateOrder"`
}

func ConvertAncientPlateResult(ancientPlate model.AncientPlate) AncientPlateResult {
	return AncientPlateResult{
		PlateId:    ancientPlate.PlateId,
		PlateName:  ancientPlate.PlateName,
		Detail:     ancientPlate.Detail,
		PlateOrder: ancientPlate.PlateOrder,
	}
}

func ConvertAncientPlateResults(ancientPlates []model.AncientPlate) []AncientPlateResult {
	results := make([]AncientPlateResult, len(ancientPlates))
	for index, value := range ancientPlates {
		results[index] = ConvertAncientPlateResult(value)
	}
	return results
}
