package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type user struct {
	Id   int    `db:"id"`
	Age  int    `db:"age"`
	Name string `db:"name"`
}

var db *sqlx.DB

func initDB(dsn string) (err error) {
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(20)
	return
}

// 查询多行数据
func queryMultiRow() {
	sqlstr := "select id, name, age from user where id > ?"
	var users []user
	err := db.Select(&users, sqlstr, 5)
	if err != nil {
		fmt.Println("select err:", err)
		return
	}
	for _, u := range users {
		fmt.Printf("id:%d name:%d age:%d\n", u.Id, u.Name, u.Age)
	}
}

// 查询单行数据
func queryRow() {
	sqlstr := "select id, name, age from user where id = ?"
	var u user
	err := db.Get(&u, sqlstr, 4)
	if err != nil {
		fmt.Println("select err:", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d", u.Id, u.Name, u.Age)
}

// 插入数据
func insertRow() {
	sqlstr := "insert into user (name, age) values (?, ?)"
	ret, err := db.Exec(sqlstr, "小王子", 13)
	if err != nil {
		fmt.Println("insert err:", err)
		return
	}
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Println("get id err:", err)
		return
	}
	fmt.Printf("insert successfully id:%d\n", id)
}

// 事务操作
func transations() {
	tx, err := db.Beginx()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		fmt.Println("Beginx err:", err)
		return
	}
	sqlstr := "update user set age=age+? where id = 5"
	tx.MustExec(sqlstr, 2) // sqlx包中的MustExec()函数优化了err处理
	sqlstr = "update user set age=age-? where id = 7"
	tx.MustExec(sqlstr, 2) // 出错后会产生panic，同时自动回滚
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		fmt.Println("Commit err:", err)
	}

}

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/sql_test"
	err := initDB(dsn)
	if err != nil {
		fmt.Println("initDB err", err)
		return
	}
	// queryMultiRow()
	// queryRow()
	// insertRow()
	transations()
}
