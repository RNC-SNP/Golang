package main

import "fmt"

func mian() {
	var array [10]int
	array[0] = 12
	array[1] = 23
	fmt.Printf("array[9]: %d", array[9])

	a1 := [3]int{1, 2, 3}
	a2 := [10]int{11, 22, 33}
	a3 := [...]int{111, 222, 333} //自动计算数组长度

	a4 := [2][3]int{[3]int{1, 2, 3}, [3]int{11, 22, 33}} //二维数组初始化
	a5 := [2][3]int{{1, 2, 3}, {11, 22, 33}}             //二维数组初始化简化（内外类型一致时）
}
