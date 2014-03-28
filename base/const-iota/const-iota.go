package main

import "fmt"

var PI float64

const (
	pi     = 3.1415926535
	prefix = "GO_"
)

const (
	a = iota //iota枚举：a=0
	b = iota //b=1
	c        //常量声明省略值时，默认和之前一个值的字面相同：c = iota 即 c = 2
)

const d = iota //每遇到一个const关键字，iota就会重置：d = 0

func main() {
	PI := 3.1415926535
	fmt.Printf("PI: %f\n", PI)
	fmt.Printf("pi: %f\n", pi)
	fmt.Printf("prefix: %s\n", prefix)
	fmt.Printf("a: %d\n", a)
	fmt.Printf("b: %d\n", b)
	fmt.Printf("c: %d\n", c)
	fmt.Printf("d: %d\n", d)
}
