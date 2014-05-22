package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func CheckIpAddress(ip string) (isValid bool) {
	if isMatch, _ := regexp.MatchString("^[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}$", ip); isMatch {
		return true
	}
	return false
}

func fetchHttpData(url string) (data string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("http get error.")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http read error")
	}

	data = string(body)
	return data
}

func filterText(txt string) {
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	txt = re.ReplaceAllStringFunc(txt, strings.ToLower)

	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	txt = re.ReplaceAllString(txt, "")

	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	txt = re.ReplaceAllString(txt, "")

	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	txt = re.ReplaceAllString(txt, "\n")

	re, _ = regexp.Compile("\\s{2,}")
	txt = re.ReplaceAllString(txt, "\n")

	fmt.Println(strings.TrimSpace(txt))
}

func main() {
	ip := "192.168.123.45"
	isIpValid := CheckIpAddress(ip)
	if isIpValid {
		fmt.Printf("%s is a valid IP address.\n", ip)
	} else {
		fmt.Printf("%s is an invalid IP address.", ip)
	}

	data := fetchHttpData("http://RincLiu.com")
	filterText(data)
}
