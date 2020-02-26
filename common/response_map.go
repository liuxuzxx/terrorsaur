package common

import (
	"github.com/kataras/iris/v12"
)

func Success(data interface{}) iris.Map {
	return iris.Map{"code": 1, "message": "success", "data": data}
}

func Failure(information interface{}) iris.Map {
	return iris.Map{"code": -1, "message": "failure", "data": information}
}
