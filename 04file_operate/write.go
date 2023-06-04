package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.OpenFile("./www.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("openFile err:", err)
		return
	}
	//file.write
	data := []byte("hello world!\n")
	n, err := file.Write(data[:])
	if err != nil {
		fmt.Println("file.write err:", err)
		return
	}
	fmt.Printf("写入 %d 个字节到文件中\n", n)

	//bufio
	//使用bufio写入文件可以提高写入效率，特别是在需要大量写入小数据块的情况下效果更为明显。
	//当然，在需要频繁写入大数据块时，使用bufio的效果并不明显。
	writer := bufio.NewWriter(file)
	defer writer.Flush() //调用Flush方法将缓冲中的数据写入文件
	n1, err := writer.Write([]byte("hello beijing!\n"))
	if err != nil {
		fmt.Println("bufio write err:", err)
		return
	}
	fmt.Printf("写入 %d 个字节到文件中\n", n1)
	file.Close()
	//ioutil
	err = ioutil.WriteFile("./www1.txt", []byte("hello china!\n"), 0644)
	if err != nil {
		fmt.Println("ioutil write err:", err)
		return
	}
	//ioutil.ReadAll、ioutil.ReadFile 和 ioutil.ReadDir等函数在1.16版本后就被弃用，可使用使用os或io包中相同的方法代替
	/*
		ioutil.ReadAll -> io.ReadAll
		ioutil.ReadFile -> os.ReadFile
		ioutil.ReadDir -> os.ReadDir
		// others
		ioutil.NopCloser -> io.NopCloser
		ioutil.ReadDir -> os.ReadDir
		ioutil.TempDir -> os.MkdirTemp
		ioutil.TempFile -> os.CreateTemp
		ioutil.WriteFile -> os.WriteFile
	*/
	err = os.WriteFile("./www1.txt", []byte("hello neimenggu!\n"), 0644)
	if err != nil {
		fmt.Println("os write err:", err)
	}

	content, err := ioutil.ReadFile("./www.txt")
	if err != nil {
		fmt.Println("read www.txt err:", err)
		return
	}
	fmt.Println("文件www.txt内容如下：")
	fmt.Println(string(content[:]))

	content1, err := ioutil.ReadFile("www1.txt")
	if err != nil {
		fmt.Println("read www1.txt err:", err)
		return
	}
	fmt.Println("文件www1.txt内容如下：")
	fmt.Println(string(content1[:]))
}
