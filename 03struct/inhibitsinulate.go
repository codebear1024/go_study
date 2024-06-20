package main

import "fmt"

// 结构体内嵌模拟 “继承”
// 结构体嵌套时：可以直接访问嵌套结构体的一个字段，使用该特性可以模拟“继承”
type animal struct {
	name string
}

// 定义一个动物会动的方法
func (a animal) move() {
	fmt.Printf("%s 会动~\n", a.name)
}

type dog struct {
	feet int
	animal
}

// 定义一个dog会叫的方法
func (d dog) wangwang() {
	fmt.Printf("%s 在叫，汪汪汪~\n", d.name)
}

func main() {
	var d = dog{
		feet: 4,
		animal: animal{
			name: "旺财",
		},
	}
	d.wangwang() //调用dog的方法
	d.move()     //调用animal的方法
}
