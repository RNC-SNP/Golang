package main

import "fmt"

func main() {
	s1 := `Hello, 
	Go lang.` //多行字符串
	fmt.Printf("s1: %s\n", s1)

	s2 := "hello, " + "go lang."
	fmt.Printf("s2: %s\n", s2)
	c := []byte(s2) //将string转换为[]byte数组
	c[0] = 'H'
	s3 := string(c) //将[]byte数组转换为string
	fmt.Printf("s3: %s\n", s3)

	s4 := "H" + s2[1:] //对string进行切片(slice)操作
	fmt.Printf("s4: %s\n", s4)
}
