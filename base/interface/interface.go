package main

import "fmt"

type Person interface { //定义interface
	SayHi()
}

type Student struct {
	name  string
	score int
}

func (s Student) SayHi() {
	fmt.Printf("Hi. My name is %s. My score is %d.\n", s.name, s.score)
}

type Teacher struct {
	name  string
	phone string
}

func (t Teacher) SayHi() {
	fmt.Printf("Hi. My name is %s. My phone is %s.\n", t.name, t.phone)
}

func main() {
	s := Student{"Rinc", 99}
	t := Teacher{"Raiser", "123456789"}
	var p Person //声明接口类型变量
	p = s        //指向接口实现类Student对象
	p.SayHi()
	p = t //指向接口实现类Teacher对象
	p.SayHi()
}
