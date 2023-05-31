package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./xxx1.txt")
	if err != nil {
		fmt.Println("open err:", err)
		return
	}
	defer file.Close()
}
