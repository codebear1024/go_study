package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Open("./xxx.txt") //windows系统下open只能打开已存在的文件，linux是可以创建文件的
	if err != nil {
		fmt.Println("open err:", err)
		return
	}
	defer file.Close()

	//file.read
	/*for {
		var buf [1024]byte
		_, err := file.Read(buf[:])
		if err == io.EOF {
			fmt.Println(string(buf[:]))
			break
		}
		if err != nil {
			fmt.Println("read err:", err)
			return
		}
		fmt.Println(string(buf[:]))
	}*/

	//bufio.read
	for {
		reader := bufio.NewReader(file)
		tmp, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println(tmp)
			break
		}
		if err != nil {
			fmt.Println("bufio.read err:", err)
			return
		}
		fmt.Print(tmp)
	}

	//ioutil
	content, err := ioutil.ReadFile("xxx1.txt")
	if err != nil {
		fmt.Println("ioutil.read err:", err)
		return
	}
	fmt.Print(string(content[:]))
}
