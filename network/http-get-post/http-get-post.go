package main
 
import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
 
func main() {
	url := "https://api.douban.com/v2/book/1220562"

	// Simple GET/POST request:
	// response, err := http.Get(url)
	// response, err := http.Post(url, "image/jpeg", &buf)
	// response, err := http.PostForm(url, url.Values{"key": {"Value"}, "id": {"123"}})
	
	// Complicated HTTP request:
	reqest, err := http.NewRequest("GET", url, nil)
	// Set headers:
	reqest.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	reqest.Header.Add("Accept-Encoding", "utf-8")
	reqest.Header.Add("Accept-Language", "zh-cn,zh;q=0.8,en-us;q=0.5,en;q=0.3")
	reqest.Header.Add("Connection", "keep-alive")
	reqest.Header.Add("Host", "api.douban.com")
	reqest.Header.Add("Referer", url)
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0")
	// Use client to send request:
	client := &http.Client{}
	response, err := client.Do(reqest)	
	
	checkErr(err)
	defer response.Body.Close()

	// Check HTTP response code:
	if response.StatusCode != 200 {
		fmt.Println("Request failed.")
		return
	}
	
	// Read HTTP response headers:
	for k, v := range response.Header {
		fmt.Println(k+":", v)
	}

	// Read HTTP response body:
	contents, err := ioutil.ReadAll(response.Body)
	checkErr(err)
	fmt.Printf("%s\n", string(contents))
}
