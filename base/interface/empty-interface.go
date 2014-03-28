package main

import (
	"fmt"
	"strconv"
)

type Element interface{} //定义空接口类型，任意类型都可被其调用

type List []Element

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return "name: " + p.name + ", age: " + strconv.Itoa(p.age)
}

func main() {
	list := make(List, 5) //创建size为3的空List数组
	list[0] = 9
	list[1] = 8.8
	list[2] = "10"
	list[3] = true
	list[4] = Person{"Rinc", 24}
	for index, element := range list {
		switch value := element.(type) { //判断类型
		case int:
			fmt.Printf("list[%d] is int: %s\n", index, value)
		case float32:
			fmt.Printf("list[%d] is float32: %s\n", index, value)
		case float64:
			fmt.Printf("list[%d] is float64: %s\n", index, value)
		case string:
			fmt.Printf("list[%d] is string: %s\n", index, value)
		case bool:
			fmt.Printf("list[%d] is bool: %s\n", index, value)
		case Person:
			fmt.Printf("list[%d] is Person: ", index)
			fmt.Println(element) //Person类型实现了Stringer接口的String()方法，在这里可直接作参数
		default:
			fmt.Printf("list[%d] is an unknown element\n", index)
		}
	}
}
