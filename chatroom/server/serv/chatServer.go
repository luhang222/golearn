package serv

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type (
	chatServer struct {
		Host  string
		Port  string
		Proto string
	}
)

func NewChatServer(host string, port string, proto string) *chatServer {
	return &chatServer{host, port, proto}
}

func (cs *chatServer) Serv() {
	addr := cs.Host + ":" + cs.Port
	listener, err := net.Listen(cs.Proto, addr)
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go readClient(conn)
	}
}

func readClient(conn net.Conn) {
	rd := bufio.NewReader(conn)
	for {
		msg, err := rd.ReadString('\n')
		msg = strings.Trim(msg, "\r\n")
		if err != nil {
			fmt.Println(err)
		}
		go spread(conn)
	}
}

func spread(conn net.Conn) {

}
