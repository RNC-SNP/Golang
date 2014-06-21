package main

import (
	"fmt"
	"net"
	"os"
	"io/ioutil"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func launchTCPClient(address string) {
	// Create TCP Address:
	tcpAddr, err := net.ResolveTCPAddr("tcp4", address)
	checkError(err)
	// Create TCP connection:
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	// Write data to server:
	_, err = conn.Write([]byte("HEAD HTTP/1.0\r\n\r\n"))
	checkError(err)
	// Read all data from server:
	result, err := ioutil.ReadAll(conn)
	checkError(err)
	// Output data from server:
	fmt.Println(string(result))
	os.Exit(0)
}

func launchUDPClient(address string) {
	// Create UDP Address:
	udpAddr, err := net.ResolveUDPAddr("udp4", address)
	checkError(err)
	// Create UDP connection:
	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)
	// Write data to server:
	_, err = conn.Write([]byte("HEAD HTTP/1.0\r\n\r\n"))
	checkError(err)
	// Read all data from server:
	result, err := ioutil.ReadAll(conn)
	checkError(err)
	// Output data from server:
	fmt.Println(string(result))
	os.Exit(0)
}

func main() {
	launchTCPClient(":8192")
	//launchUDPClient(":8192")
}
