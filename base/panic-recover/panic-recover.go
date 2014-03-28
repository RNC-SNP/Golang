package main

import (
	"fmt"
	"os"
)

var user string //声明变量

/*
init()与main()比较:
1. 都是保留函数，都是自动调用的;
2. init()一个package可有多个, main()全局只能有一个;
3. init()可用于所有package, main()只能用于main package;
4. init()是可选的, main()是必须的;
*/
func init() {
	user = os.Getenv("USER") //读取系统环境变量
}

func printUser() {
	if user == "" {
		panic("$USER is NIL !!!") //抛出异常信息
	} else {
		fmt.Println("$USER:", user)
	}
}

func recoverPanic(f func()) { //所有类型函数作为参数传入
	defer func() { //recover仅在defer函数中有效
		if errorMessage := recover(); errorMessage != nil { //调用recover得到panic时的错误信息
			fmt.Println("Recovered from a panic with error message:'", errorMessage, "'")
		}
	}()
	f() //执行函数f，如果f中出现了panic，那么就可以恢复回来
}

func main() {
	recoverPanic(printUser)
}
