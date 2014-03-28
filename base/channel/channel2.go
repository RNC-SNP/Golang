package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	isTimeout := make(chan bool)
	go func() {
		for {
			select {
			case v := <-ch:
				fmt.Println(v)
			case <-time.After(5 * time.Second): //设置五秒超时
				fmt.Println("timeout...")
				isTimeout <- true
				break
			}
		}
	}()
	<-isTimeout
}
