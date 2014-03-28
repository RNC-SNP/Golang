package main

import (
	"fmt"
)

func sum(array []int, ch chan int) {
	result := 0
	for _, value := range array {
		result += value
	}
	ch <- result //结果发送到管道ch
}

func main() {
	array := []int{1, 3, 5, 7, 9}
	ch1 := make(chan int)             //创建默认无缓冲区的管道
	go sum(array[:len(array)/2], ch1) //并发执行
	go sum(array[len(array)/2:], ch1)
	result1, result2 := <-ch1, <-ch1 //从管道接收数据
	fmt.Println(result1, result2)

	size := 10
	ch2 := make(chan int, size) //创建指定大小缓冲区的管道
	for i := 0; i < size; i++ {
		ch2 <- i
	}
	close(ch2)          //关闭管道
	v, isClose := <-ch2 //测试管道是否关闭
	fmt.Println("Is the channel closed:", v, isClose)
	for result := range ch2 { //管道关闭后必须通过range逐个读取数据
		fmt.Println(result)
	}
}
