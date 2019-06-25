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
}

func swap(left, right string) (string, string, string) {
	return right, left, left + right
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
