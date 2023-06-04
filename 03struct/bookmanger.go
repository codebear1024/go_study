package main

import "fmt"

/*
需求：
1.实现一个简单的图书管理系统
2.每本书有书名、作者、价格、上架信息
3.用户可以在控制台添加和删除书籍，可以修改书籍信息，显示所有书籍信息
*/
type book struct {
	name   string
	author string
	price  float64
	status bool
}

var bookslice []book

func main() {
	for {
		var cmd int
		showmenu()
		fmt.Scanln(&cmd)
		switch cmd {
		case 1:
			addormodifybook(cmd)
		case 2:
			addormodifybook(cmd)
		case 3:
			showbook()
		case 4:
			var bookName string
			fmt.Print("请输入要删除的书名：")
			fmt.Scanln(&bookName)
			delbook(bookName)
		case 5:
			return
		}
	}
}

func showmenu() {
	fmt.Println("1. 添加书籍")
	fmt.Println("2. 修改数据信息")
	fmt.Println("3. 展示所有书籍")
	fmt.Println("4. 删除书籍")
	fmt.Println("5. 退出程序")
}

func addBook(bookname, author string, price float64, status bool) {
	bookslice = append(bookslice, book{bookname, author, price, status})
}

func modifyBook(bookname, author string, price float64, status bool) {
	for i := 0; i < len(bookslice); i++ {
		if bookslice[i].name == bookname {
			bookslice[i].author = author
			bookslice[i].price = price
			bookslice[i].status = status
		}
	}
}

func showbook() {
	for i := 0; i < len(bookslice); i++ {
		fmt.Printf("name:%s, author:%s, price:%v, status:%v\n", bookslice[i].name, bookslice[i].author, bookslice[i].price, bookslice[i].status)
	}
}

func delbook(bookname string) {
	for i := 0; i < len(bookslice); i++ {
		if bookslice[i].name == bookname {
			copy(bookslice[i:], bookslice[i+1:])
			bookslice = bookslice[:len(bookslice)-1]
		}
	}
}

func addormodifybook(cmd int) {
	fmt.Println("请依次按提示输入相关信息。")
	var bookName string
	var author string
	var price float64
	var status bool
	fmt.Print("name:")
	fmt.Scanln(&bookName)
	fmt.Print("author:")
	fmt.Scanln(&author)
	fmt.Print("price:")
	fmt.Scanln(&price)
	fmt.Print("status:")
	fmt.Scanln(&status)
	switch cmd {
	case 1:
		addBook(bookName, author, price, status)
	case 2:
		modifyBook(bookName, author, price, status)
	}
}
