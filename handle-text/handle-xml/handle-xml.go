package main

import (
	"fmt"
	"encoding/xml"
	"io/ioutil"
	"os"
)

// Define struct type with reflection tags to describe the xml doc's structure.
// All of the properties' name should start with a uppercase character to make sure they are accessible.

type ViewGroup struct {
	// Declare element's name:
	XMLName xml.Name `xml:"ViewGroup"`
	// Declare element's attribute:
	IsVisible string `xml:"IsVisible,attr"`
	// Declare element's children:
	Views []View `xml:"View"`
	// Append its inner xml:
	//Description string `xml:",innerxml"`
}

type View struct {
	//XMLName xml.Name `xml:"View"`
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

	// Parse xml data to object:
	viewGroup1 := ViewGroup{}
	err = xml.Unmarshal(data, &viewGroup1)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Println(viewGroup1)

	// Parse object to xml data:
	viewGroup2 := &ViewGroup{IsVisible: "false"}
	viewGroup2.Views = append(viewGroup2.Views, View{"12dp", "34dp"})
	viewGroup2.Views = append(viewGroup2.Views, View{"56dp", "78dp"})
	output, err := xml.MarshalIndent(viewGroup2, "", "")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	// Write default xml header:
	//os.Stdout.Write([]byte{xml.Header})
	os.Stdout.Write(output)
}
