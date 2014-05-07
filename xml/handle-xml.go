package main

import (
	"fmt"
	"encoding/xml"
	"io/ioutil"
	"os"
)

// Define struct types to describe xml doc's structure:

type ViewGroup struct {
	XMLName xml.Name `xml:"ViewGroup"`
	IsVisible string `xml:"IsVisible,attr"`
	Views []View `xml:"View"`
	Description string `xml:",innerxml"`
}

type View struct {
	XMLName xml.Name `xml:"View"`
	Width string `xml:"Width"`
	Height string `xml:"Height"`
}

func main() {
	// Open xml file:
	file, err := os.Open("layout.xml")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()

	// Read xml file:
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	// Parse xml data:
	v := ViewGroup{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Println(v)
}
