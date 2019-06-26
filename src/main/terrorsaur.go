package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("尝试使用一下Idea进行Go语言的开发")
	first, second, third := swap("liuxu", "zhongxiaoxia")
	fmt.Println(first, second, third)
	fmt.Println("查看获取到的数据信息:", math.Pi)
	fmt.Println(split(17))
	fmt.Println(sum(34, 89))
	fmt.Println(sumString("左手", "右手"))

	logVarInformation()
	replaceVar()
	showBasicTypes()
	convertType()
}

/**
 * Go的函数奇怪啊，这种类型放在名字后面有点美国人的味道
 */
func sum(left int, right int) int {
	return left + right
}

func sumString(left, right string) string {
	return left + right
}

func swap(left, right string) (string, string, string) {
	return right, left, left + right
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

/**
实验var关键字
感觉这个var就是声明一下，我这一行是变量variables
总体的方式就是: var 变量名字 变量类型 = 初始化的数值
*/
var language, c, cpp, python, java, golang = "language", "C", "C++", "Python", "Java", "Golang"

func logVarInformation() {
	var count, numberStr = 1, "Yes"
	fmt.Println(count, language, c, cpp, python, java, golang, numberStr)
}

/**
:= 运算符号的使用说明 说白了就是替换var
但是在函数外面不能使用
使用这个可以推断inffered 类型
*/
func replaceVar() {
	userName, password := "root", "root123"
	fmt.Println(userName, password)
}

/**
Golang的变量系列
*/
var (
	flag      bool   = true
	cluesName string = "线索的名字"
	count     uint64 = 1<<64 - 1
)

func showBasicTypes() {
	fmt.Println(flag, cluesName, count)
}

func convertType() {
	const goLangVersion = "1.12.6"
	driverClassName, connectionCount := "com.mysql.jdbc.Driver", 60
	fmt.Print(driverClassName, float64(connectionCount), "Golang的版本号:", goLangVersion)
}

/**
感觉GO的所有只要是你想声明多个的就使用括号就行了
*/
const (
	driverClassName = "com.mysql.jdbc.Driver"
	url             = "jdbc:mysql://localhost:3306"
	port            = 3306
	userName        = "root"
	password        = "root124"
)
