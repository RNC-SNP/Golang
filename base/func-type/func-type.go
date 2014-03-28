package main

import "fmt"

//定义一种函数类型IntBool
type IntBool func(int) bool //所有带一个int参数、返回一个bool值的函数都属于这种类型

func isUnsigned(x int) bool {
	if x >= 0 {
		return true
	}
	return false
}

func filter(array []int, intBool IntBool) []int { //使用函数类型作参数类型
	var result []int
	for _, value := range array {
		if intBool(value) {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	array := []int{-1, 2, 0, -3, 4, 5}
	fmt.Println("array:", array)
	array1 := filter(array, isUnsigned) //函数作为参数值
	fmt.Println("unsigned:", array1)
}
