package main

import "fmt"

func main() {
	if x := 9; x > 10 { //if条件内可声明局部变量，但仅限于if-else代码块内部
		fmt.Println(x, " is greater than 10.")
	} else {
		fmt.Println(x, "x is less than 10.")
	}

	sum := 1
	for i := 1; i <= 15; i++ { //当for只有中间的条件时，甚至两个分号都能省略，这时相当于C语言的while语句
		sum *= i
	}
	fmt.Println("15! =", sum)

	array := [...]string{"A", "B", "C", "D", "E"}
	array1 := array[1:4]
	for x := range array1 { //for配合range从array读取数据
		fmt.Println("array1[", x, "]:", array1[x])
	}

	mp := map[string]float32{"C": 5, "GO": 4.5, "JAVA": 6}
	for key, value := range mp { //for配合range从map读取数据
		fmt.Println(key, ":", value)
	}
	for key, _ := range mp { //用"_"丢弃不用的返回值
		fmt.Println("key:", key)
	}
	for _, value := range mp { //用"_"丢弃不用的返回值
		fmt.Println("value:", value)
	}

	x := 10
	switch x {
	case 0:
		fmt.Println("x is 0") //默认不会贯穿，不需要break
	case 1, 5, 10: //支持合并多个case
		fmt.Println("x is greater than 0")
		fallthrough //使用fallthrough关键字贯穿
	default:
		fmt.Println("default")
	}
}
