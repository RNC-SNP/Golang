package main

import (
	"fmt"
	"html/template"
	"net/http"
	"log"
)

func welcome(writer http.ResponseWriter, request *http.Request){
	if request.Method == "GET" {
		// Load template file:
		template, _ := template.ParseFiles("welcome.gtpl")
		template.Execute(writer, nil)
	} else {
		fmt.Fprintf(writer, "Invalid request!")
	}
}

func login(writer http.ResponseWriter, request *http.Request){
	if request.Method == "GET" {
		// Load template file:
		template, _ := template.ParseFiles("login.gtpl")
		template.Execute(writer, nil)
	} else {
		// Handle post request:
		request.ParseForm()
		fmt.Println(request.Form)
		user := request.Form["username"]
		psw := request.Form["password"]
		if  user != nil && psw !=nil {
			fmt.Fprintf(writer, "Welcome, %s", user)
		} else {
			fmt.Fprintf(writer, "Login failed!")
		}
	}
}

func main(){
	http.HandleFunc("/", welcome)
	http.HandleFunc("/login/", login)
	error := http.ListenAndServe(":12345", nil)
	if error != nil {
		log.Fatal("Server error:", error)
	}
}
