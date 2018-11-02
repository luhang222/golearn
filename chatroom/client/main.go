package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const HOST = "localhost"
const PORT = "52535"

func main() {
	fmt.Println("Starting Client")
	conn, err := net.Dial("tcp", HOST+":"+PORT)
	if err != nil {
		fmt.Println("Error Dialing", err.Error())
	}
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("welcome to Chatroom, what's your name?")
	clientName, err := inputReader.ReadString('\n')
	clientName = strings.Trim(clientName, "\r\n")

	for {
		msg, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("error input")
		}
		msg = strings.Trim(msg, "\r\n")
		if msg == "exit" {
			return
		}
		_, err = conn.Write([]byte(msg))
	}
}
