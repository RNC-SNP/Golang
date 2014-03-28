package main

import "fmt"

type Person struct { //定义struct类型
	name string
	age  int
}

type Student struct {
	Person //如果匿名成员也是struct类型，则包含其所有属性，相当于继承
	score  int
}

func compareAge(p1, p2 Person) (Person, int) {
	if p1.age >= p2.age {
		return p1, p1.age - p2.age
	} else {
		return p2, p2.age - p1.age
	}
}

func compareScore(s1, s2 Student) (Student, int) {
	if s1.score >= s2.score {
		return s1, s1.score - s2.score
	} else {
		return s2, s2.score - s1.score
	}
}

func main() {
	var p1 Person
	p1.name, p1.age = "Rinc", 24
	p2 := Person{name: "Raiser", age: 25}
	s1 := Student{Person: p1, score: 98}
	s2 := Student{Person: p2, score: 89}
	p, a := compareAge(s1.Person, s2.Person) //匿名成员作参数
	fmt.Println(p.name, "is older:", a)
	st, sc := compareScore(s1, s2)
	fmt.Println(st.Person.name, "'s score is higher:", sc)
}
