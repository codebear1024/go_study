package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

type client struct {
	ch         chan string
	lastRecord time.Time
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

/* 当与n个客户端保持聊天session时，这个程序会有2n+2个并发的goroutine
 * 这个程序却并不需要显式的锁，clients这个map被限制在了一个独立的goroutine中，所以它不能被并发地访问。
 * 多个goroutine共享的变量只有这些channel和net.Conn的实例，两个东西都是并发安全的。
 */
func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("Listen err:", err)
		return
	}
	defer listener.Close()
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept err:", err)
		}
		go handleConn(conn)
	}
}

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli.ch <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}

func handleConn(conn net.Conn) {
	cli := client{make(chan string), time.Now()}
	go clientWrite(conn, cli)
	who := conn.RemoteAddr().String()
	cli.ch <- "You are " + who
	messages <- who + " has arrived!"
	entering <- cli
	input := bufio.NewScanner(conn)
	for {
		ok := input.Scan()
		if !ok {
			return
		}
		messages <- who + "：" + input.Text()
		cli.lastRecord = time.Now()
	}
	leaving <- cli
	messages <- who + " has left!"
}

func clientWrite(conn net.Conn, cli client) {
	/*for msg := range ch {
		fmt.Fprintln(conn, msg)
	}*/
	for {
		select {
		case msg := <-cli.ch:
			fmt.Fprintln(conn, msg)
		default:
			if time.Now().Sub(cli.lastRecord) > 20*time.Second {
				leaving <- cli
				who := conn.RemoteAddr().String()
				messages <- who + " has be closed!"
				conn.Close()
				fmt.Printf("%s has been closed=============\n", who)
				return
			}
		}
	}
}
