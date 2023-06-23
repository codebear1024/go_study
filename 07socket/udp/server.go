package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	listener, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 1314,
	})
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}
	defer listener.Close()
	for {
		var buf [1024]byte
		n, client_addr, _ := listener.ReadFromUDP(buf[:])
		if n > 0 {
			fmt.Printf("%v: %v", client_addr.String(), string(buf[:n]))
			reader := bufio.NewReader(os.Stdin)
			message, _ := reader.ReadString('\n')
			listener.WriteToUDP([]byte(message), client_addr)
		}
	}
}
