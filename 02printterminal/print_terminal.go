package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 使用 fmt.Scanln() 方式读取输入
	var name string
	fmt.Print("请输入名字：")
	fmt.Scanln(&name)
	fmt.Printf("你的名字是 %s\n", name)

	// 使用 bufio.NewReader().ReadString() 方式读取输入
	fmt.Print("请输入一句话：")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Printf("你刚才说的是：%s", text)

	// 使用 os.Stdin.Read() 方式读取输入
	fmt.Print("请输入一句话：")
	data := make([]byte, 1024)
	length, _ := os.Stdin.Read(data)
	fmt.Printf("你刚才说的是：%s", string(data[:length]))

	// 不断读取输入直到用户输入回车
	fmt.Println("请一句一句说话，回车结束：")
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		if len(text) == 1 { // 用户只输入了回车
			fmt.Println("听不懂你在说什么")
		} else {
			fmt.Printf("你刚才说的是：%s", text)
		}
	}
}
