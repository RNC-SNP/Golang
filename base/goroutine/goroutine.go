package main

import (
	"fmt"
	"runtime"
)

func say(str string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(str)
		runtime.Gosched() //让出CPU时间片
	}
}

func main() {
	//runtime.GOMAXPROCS(20) //指定最大并发数
	go say("golang", 10) //开启一个新的goroutine执行
	say("hello", 10)
}
