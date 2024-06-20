package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type book struct {
	Id    int
	Title string
	Price float64
}

func initDB(dsn string) (err error) {
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(16)
	return
}

func insertBook(title string, price int) (id int64, err error) {
	sqlstr := "insert into book(title,price) values(?,?)"
	ret, err := db.Exec(sqlstr, title, price)
	id, err = ret.LastInsertId()
	return
}

func listBook() (books []book, err error) {
	sqlstr := "select id, title, price from book where id > 0"
	rows, err := db.Query(sqlstr)
	if err != nil {
		return
	}
	for rows.Next() {
		var book book
		err = rows.Scan(&book.Id, &book.Title, &book.Price)
		if err != nil {
			continue
		}
		books = append(books, book)
	}
	return
}

func deleteBook(id int) (err error) {
	sqlstr := "delete from book where id = ?"
	_, err = db.Exec(sqlstr, id)
	return
}

func modifyBook(id int, price int) {

}
