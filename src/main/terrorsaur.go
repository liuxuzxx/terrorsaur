package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
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
	equalDifferenceSeries(100)
	binaryInt(20000)
	squrt(-100)
	deferTime()
	checkSystemVersion()
	pointer()
	vertexInformation()
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

/**
进入Go的控制类型的语句试验场
发现Go声明变量怪怪的
要么就是
var 变量名 类型 = 初始化的数据
要么就是 变量名 := 初始化的数据
*/
func equalDifferenceSeries(maxNumber int) {
	var sum = 0
	for index := 0; index < maxNumber; index++ {
		sum += index
	}
	fmt.Println(maxNumber, "等差数列的和是:", sum)
}

func binaryInt(maxNumber int) {
	sum := 1
	for sum < maxNumber {
		sum += sum
	}
	fmt.Println("获取到的最大二进制:{}", sum)
}

func squrt(number float64) {
	if number < 0 {
		fmt.Println("小于0的数字没法操作")
	} else {
		fmt.Println("开平发的数字是:", math.Sqrt(number))
	}
}

func checkSystemVersion() {
	fmt.Print("Go语言运行的操作系统环境")
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Linux操作系统")
	case "windows":
		fmt.Println("Windows操作系统:", os)
	default:
		fmt.Print("不知道啥子系统:", os)
	}

	fmt.Println(time.Thursday)
}

/**
defer这个关键字怎么看怎么像 try{函数所有的代码}finally{defer语句}
*/
func deferTime() {
	defer fmt.Println("推迟执行")
	fmt.Println("在这之前执行")
}

/**
坏了，Go越看越像c了，简直就是一个翻版的c啊，指针来了，其他还会远吗
看看指针
*/

func pointer() {
	number := 90
	numberPoint := &number

	fmt.Println(numberPoint)
}

type Vertex struct {
	length int
	width  int
	height int
}

func vertexInformation() {
	fmt.Print(Vertex{10, 90, 89})
}
