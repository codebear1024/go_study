
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("server start")
	listen, err := net.Listen("tcp", "127.0.0.1:1213")
	if err != nil {
		fmt.Println("linsten err:", err)
		return
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept err:", err)
			return
		}
		process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	go recv_message(conn)
	for {
		reader := bufio.NewReader(os.Stdin)
		message, _ := reader.ReadString('\n')
		conn.Write([]byte(message))
	}
}

func recv_message(conn net.Conn) {
	for {
		var buf [1024]byte
		n, _ := conn.Read(buf[:])
		if n > 0 {
			fmt.Print(string(buf[:n]))
		}
	}
}
