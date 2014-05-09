package main

import (
	"fmt"
	"net"
	"os"
	"time"
	"io/ioutil"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func launchTCPServer(address string) {
	// Create TCP Address:
	tcpAddr, err := net.ResolveTCPAddr("tcp4", address)
	checkError(err)
	// Listen TCP address:
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	// Accept the clients' connection to handle:
	for {
		conn, err := listener.Accept()
		// When an error occured, the server should continue to handle other client:
		if err != nil {
			continue
		}
		// Create new thread to handle TCP client:
		go handleTCPClient(conn)
	}
}

func handleTCPClient(conn net.Conn) {
	defer conn.Close()
	// Write data to client:
	datetime := time.Now().String()
	conn.Write([]byte(datetime))
	// Read all data from client:
	result, err := ioutil.ReadAll(conn)
	checkError(err)
	// Output data from client:
	fmt.Println(string(result))
}

func launchUDPServer(address string) {
	// Create UDP Address:
	udpAddr, err := net.ResolveUDPAddr("udp4", address)
	checkError(err)
	// Listen UDP address:
	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)
	// Accept the clients' connection to handle:
	for {
		handleUDPClient(conn)
	}
}

func handleUDPClient(conn *net.UDPConn) {
	// Read address from client:
	var buf [512]byte
	_, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}
	// Write data to client:
	datetime := time.Now().String()
	conn.WriteToUDP([]byte(datetime, addr)
}

func main() {
	launchTCPServer(":8192")
	//launchUDPServer(":8192")
}
