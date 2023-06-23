package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:1213")
	if err != nil {
		fmt.Println("dial err:", err)
		return
	}
	defer conn.Close()
	go client_recv_message(conn)
	for {
		reader := bufio.NewReader(os.Stdin)
		message, _ := reader.ReadString('\n')
		conn.Write([]byte(message))
	}
}

func client_recv_message(conn net.Conn) {
	for {
		var buf [1024]byte
		n, _ := conn.Read(buf[:])
		if n > 0 {
			fmt.Print(string(buf[:n]))
		}
	}
}
