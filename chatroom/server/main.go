package main

import (
	"fmt"
	"net"
)

const HOST = "localhost"
const PORT = "52535"

func main() {
	fmt.Println("Starting Server")
	listener, err := net.Listen("tcp", HOST+":"+PORT)
	if err != nil {
		fmt.Println("Error Listening", err.Error())
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error Accepting", err.Error())
			panic(err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	for {
		buf := make([]byte, 1024)
		byteNum, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
		}
		fmt.Printf("Receivng data: %d,%v \r\n", byteNum, string(buf))
	}
}
