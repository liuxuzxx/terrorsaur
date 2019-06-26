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
