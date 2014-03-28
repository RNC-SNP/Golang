package main

import "fmt"

//普通函数
func factorial(x int) int { //返回值变量名可省略
	result := 1
	for i := 2; i <= x; i++ {
		result *= i
	}
	return result
}

//指针函数
func factorialP(x *int) int {
	result := 1
	for i := 2; i <= *x; i++ {
		result *= i
	}
	return result
}

//多返回值函数
func calc(x, y float64) (float64, float64, float64, float64) {
	return x + y, x - y, x * y, x / y
}

//不定参数
func sum(args ...int) int {
	result := 0
	for argIndex, argValue := range args { //从array读取参数
		fmt.Println("arg", argIndex, ":", argValue)
		result += argValue
	}
	return result
}

func main() {
	x := 15
	fmt.Println(x, "! =", factorial(x))
	fmt.Println(x, "! =", factorialP(&x))
	res1, res2, res3, res4 := calc(8, 9)
	fmt.Println("8 + 9=", res1)
	fmt.Println("8 - 9=", res2)
	fmt.Println("8 * 9=", res3)
	fmt.Println("8 / 9=", res4)
	fmt.Println("sum =", sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	for i := 1; i < 10; i++ {
		defer fmt.Println("i :", i) //defer关键字用于延迟执行，采取先进后出模式
	}
}
