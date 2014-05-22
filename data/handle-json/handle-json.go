package main

import (
	"encoding/json"
	"fmt"
)

// Define struct type with reflection tags to describe the json doc's structure.
// All of the properties' name should start with a uppercase character to make sure they are accessible.

type View struct {
	Width string `json:"width"`
	Height string `json:"height"`
}

type ViewGroup struct {
	Views []View `json:"views"`
}

func main() {
	// Parse string to JSON object:
	var viewGroup1 ViewGroup
	str := `{"Views":[{"Width":"56dp","Height":"78dp"},{"Width":"12dp","Height":"34dp"}]}`
	json.Unmarshal([]byte(str), &viewGroup1)
	fmt.Println(viewGroup1)

	// Parse JSON object to string:
	var viewGroup2 ViewGroup
	viewGroup2.Views = append(viewGroup2.Views, View{Width:"120dp", Height:"340dp"})
	viewGroup2.Views = append(viewGroup2.Views, View{Width:"560dp", Height:"780dp"})
	bytes, err := json.Marshal(viewGroup2)
	if err != nil {
		fmt.Printf("error:", err)
	}
	fmt.Println(string(bytes))
}
