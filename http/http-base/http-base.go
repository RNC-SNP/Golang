package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHello(responseWriter http.ResponseWriter, request *http.Request) {
	request.ParseForm()       //解析表单参数，默认不会解析
	fmt.Println(request.Form) //打印请求的相关信息
	fmt.Println("path:", request.URL.Path)
	fmt.Println("UserAgent:", request.UserAgent())
	for k, v := range request.Form {
		fmt.Printf("key: %s, value: %s.\n", k, v)
	}
	fmt.Fprintf(responseWriter, "Hello, Golang.") //写入writer
}

type MyHttpHandler struct{} //定义一个实现了Http.Handler接口的类型

func (handler *MyHttpHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	if request.URL.Path == "/" {
		sayHello(responseWriter, request)
	} else {
		http.NotFound(responseWriter, request)
	}
}

func main() {
	http.HandleFunc("/", sayHello)              //设置访问的路由:第一个参数为路径，第二个参数是对应的处理函数
	error := http.ListenAndServe(":18188", nil) //设置端口监听,返回错误信息
	if error != nil {
		log.Fatal("ListenAndServe: ", error)
	}

	myHttpHandler := &MyHttpHandler{}
	http.ListenAndServe(":19199", myHttpHandler) //调用自定义HttpHandler
}
