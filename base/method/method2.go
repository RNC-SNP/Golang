package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h Human) SayHi() {
	fmt.Printf("Hi. My name is %s. I'm %d years old.\n", h.name, h.age)
}

type Student struct {
	Human //Human的sayHi方法也被继承
	score int
}

func (s Student) SayHi() { //重写父类method
	fmt.Printf("Hi. My name is %s. I'm %d years old. My score is %d.\n", s.Human.name, s.Human.age, s.score)
}

func main() {
	p := Human{"Rinc", 24}
	s := Student{p, 99}
	s.Human.SayHi() //调用父类method
	s.SayHi()       //调用重写的method
}
