package main

import "fmt"

func fibonacci(ch, quit chan int) {
	x, y := 1, 1
	for {
		/*
			select用于多个channel,随机选择一个可用的管道
		*/
		select {
		case ch <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
			//default:
			//	fmt.Println("blocking...")
		}
	}
}

func main() {
	ch := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}
		quit <- 0
	}()
	fibonacci(ch, quit)
}
