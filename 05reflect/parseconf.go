package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 写一个程序，解析配置文件
type configMysql struct {
	username string
	password string
	addr     string
	port     uint32
}

type configRedis struct {
	host string
	port uint32
}

func main() {
	// 打开文件
	confile, err := os.Open("./xxx.ini")
	if err != nil {
		fmt.Println("open", err)
		return
	}
	defer confile.Close()
	// 一行一行的读文件
	reader := bufio.NewReader(confile)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if line == "\r\n" {
			continue
		}
		// 用strings.TrimSpace()删除开头和结尾的空格
		//line = strings.TrimSpace(line)
		// 用string。ReplaceAll()删除中间的空格
		line = strings.ReplaceAll(line, " ", "")
		fmt.Printf(line)
	}
	// 解析每一行读出来的文件
	// 按照‘=’前的名字找到对应结构体字段，并赋值
	// 将结构体按一定的格式打印出来
}
