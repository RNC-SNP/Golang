package main

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"crypto/md5"
	"io"
	"time"
	"strconv"
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

func register(writer http.ResponseWriter, request *http.Request){
	if request.Method == "GET" {
		// Read system timestamp to generate token:
		current := time.Now().Unix()
		_md5 := md5.New()
		io.WriteString(_md5, strconv.FormatInt(current, 10))
		token := fmt.Sprintf("%x", _md5.Sum(nil))		

		// Load template file and send token to client:
		template, _ := template.ParseFiles("register.gtpl")
		template.Execute(writer, token)
	} else {
		// Parse form fields:
		request.ParseForm()
		fmt.Println(request.Form)
		formTokenArray := request.Form["token"]
		formNameArray := request.Form["name"]
		formGenderArray := request.Form["gender"]
		formLanguageArray := request.Form["language"]
		formPlatformArray := request.Form["platform"]	

		// Check token:
		if formTokenArray[0] == "" {
			fmt.Fprintf(writer, "No token!\n")
		} else {
			fmt.Fprintf(writer, "Your token: %s\n", formTokenArray[0])
		}

		// Check text length and format:
		if  len(formNameArray[0]) <= 0 {
			fmt.Fprintf(writer, "Name length too short: %s\n", formNameArray[0])
		} else if isMatch, _ := regexp.MatchString("^[a-zA-Z]+$", formNameArray[0]); !isMatch {
			fmt.Fprintf(writer, "Name format is invalid: %s\n", formNameArray[0])
		} else {
			fmt.Fprintf(writer, "Your name: %s\n", formNameArray[0])
		}

		// Check radio:
		if values := []string{"Male", "Female"}; !checkField(formGenderArray[0], values) {
			fmt.Fprintf(writer, "Gender is invalid: %s\n", formGenderArray[0])
		} else {
			fmt.Fprintf(writer, "Your gender: %s\n", formGenderArray[0])
		}	

		fmt.Fprintf(writer, "Your language: %s\n", formLanguageArray)
		fmt.Fprintf(writer, "Your platform: %s\n", formPlatformArray)			
	}
}

func checkField(field string, values []string)bool {
	for _, v := range values {
		if v == field {
			return true
		}
	}
	return false
}

func main(){
	http.HandleFunc("/", welcome)
	http.HandleFunc("/register/", register)
	error := http.ListenAndServe(":12345", nil)
	if error != nil {
		log.Fatal("Server error:", error)
	}
}
