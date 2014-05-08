package main

import (
	"fmt"
	"os"
)

func operateDirectory() {
	// Make directories:
	os.Mkdir("folder-a", 0777)
	os.MkdirAll("folder-b/folder-ba", 0777)

	// Remove directories:
	err := os.Remove("folder-a")
	if err != nil {
		fmt.Println(err)
	}
	err = os.RemoveAll("folder-b")
	if err != nil {
		fmt.Println(err)
	}
}

func writeFile(fileStr string) {
	// Create file:
	fout, err := os.Create(fileStr)
	if err != nil {
		fmt.Println(fileStr, err)
		return
	}
	// Write string and byte array to file:
	for i := 0; i < (1 << 10); i++ {
		fout.WriteString("write a string...\n")
		fout.Write([]byte("write a byte array...\n"))
	}
}

func readFile(fileStr string) {
	// Open file:
	file, err := os.Open(fileStr)
	if err != nil {
		fmt.Println(fileStr, err)
	}
	// Create a byte array to buffer data:
	buffer := make([]byte, 1024)
	// Read data to buffer and output:
	for {
		length, _ := file.Read(buffer)
		if length == 0 {
			break;
		}
		os.Stdout.Write(buffer[:length])
	}
}

func deleteFile(fileStr string) {
	err := os.Remove(fileStr)
	if err != nil {
		fmt.Println(fileStr, err)
	}
}

func main() {
	operateDirectory()
	fileStr := "txt.txt"	
	writeFile(fileStr)	
	readFile(fileStr)	
	deleteFile(fileStr)
}
