package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func initTcpServer(address string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", address)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleTcpClient(conn)
	}
}

func handleTcpClient(conn net.Conn) {
	defer conn.Close()
	datetime := time.Now().String()
	conn.Write([]byte(datetime))
}

func main() {
	initTcpServer(":8192")
}
