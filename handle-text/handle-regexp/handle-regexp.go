package main

import (
	"fmt"
	//"io/ioutil"
	//"net/http"
	"regexp"
	//"strings"
)

func CheckIpAddress(ip string) (isValid bool) {
	if isMatch, _ := regexp.MatchString("^[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}$", ip); isMatch {
		return true
	}
	return false
}

func main() {
	ip := "192.168.123.45"
	isIpValid := CheckIpAddress(ip)
	if isIpValid {
		fmt.Printf("%s is a valid IP address.\n", ip)
	} else {
		fmt.Printf("%s is an invalid IP address.", ip)
	}
}
