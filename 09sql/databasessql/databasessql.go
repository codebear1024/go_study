package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	id   int
	age  int
	name string
}

var DB *sql.DB

func initDB(dsn string) (err error) {
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	err = DB.Ping()
	if err != nil {
		return
	}
	DB.SetMaxOpenConns(50)
	DB.SetMaxIdleConns(10)
	return
}

// 查询单行数据
func queryRowDemo(id int) {
	sqlstr := "select id, name, age from user where id=?"
	var u user
	err := DB.QueryRow(sqlstr, id).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Println("Scan err:", err)
		return
	}
	fmt.Printf("query successfully, id: %d, name: %s, age: %d", u.id, u.name, u.age)
}

// 查询多行数据
func queryMultiRowDemo() {
	sqlstr := "select id, name, age from user where id >?"
	rows, err := DB.Query(sqlstr, 0)
	if err != nil {
		fmt.Println("query err:", err)
		return
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()
	for rows.Next() {
		var u user
		err = rows.Scan(&u.id, &u.name, &u.age) //如果查询后不使用scan取查询结果，等达到最大连接数就会阻塞
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}

// 插入单行数据
func insertRowDemo() {
	sqlstr := "insert into user(name,age) values(?,?)"
	ret, err := DB.Exec(sqlstr, "zhaoliu", 32)
	if err != nil {
		fmt.Println("Inserting err:", err)
		return
	}
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Println("get lastInsertId err:", err)
		return
	}
	fmt.Println("Insert successed, the id is", id)
}

// 删除数据
func deleteRowDemo(id int) {
	sqlstr := "delete from user where id = ?"
	ret, err := DB.Exec(sqlstr, id)
	if err != nil {
		fmt.Println("delete err:", err)
		return
	}
	n, err := ret.RowsAffected() //操作影响的行数
	if err != nil {
		fmt.Println("RowsAffected err:", err)
		return
	}
	fmt.Println("delete successfully, affected rows is:", n)
}

// 更新数据
func updateRowDemo(age, id int) {
	sqlstr := "update user set age=? where id=?"
	ret, err := DB.Exec(sqlstr, age, id)
	if err != nil {
		fmt.Println("update err:", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Println("RowsAffected err:", err)
		return
	}
	fmt.Println("update successfully, affected rows is:", n)
}

// stmt预处理查询
func prepareQueryDemo() {
	sqlstr := "select id, name, age from user where id=?"
	stmt, err := DB.Prepare(sqlstr)
	if err != nil {
		fmt.Println("prepare err:", err)
		return
	}
	defer stmt.Close()
	for i := 0; i < 5; i++ {
		var (
			id   int
			name string
			age  int
		)
		err = stmt.QueryRow(i).Scan(&id, &name, &age)
		if err != nil {
			fmt.Println("query err:", err.Error())
			continue
		}
		fmt.Printf("query successfully id: %d, name: %s, age: %d\n", id, name, age)
	}
}

// 预处理插入,更新和删除与插入预处理相似
func prepareInsertDemo() {
	sqlstr := "insert into user (name, age) values (?,?)"
	stmt, err := DB.Prepare(sqlstr)
	if err != nil {
		fmt.Println("Prepare err:", err)
		return
	}
	defer stmt.Close()
	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("user_%02d", i)
		ret, err := stmt.Exec(name, 20+i)
		if err != nil {
			fmt.Println("insert err:", err.Error())
			continue
		}
		id, err := ret.LastInsertId()
		fmt.Println("insert successfully, id is", id)
	}

}

// 事务操作
func transationDemo() {
	tx, err := DB.Begin() // 开启事务
	if err != nil {
		if tx == nil {
			_ = tx.Rollback() // 回滚
		}
		fmt.Println("Begin err:", err)
		return
	}
	sqlstr := "update user set age=age+? where id=?"
	_, err = tx.Exec(sqlstr, 2, 5)
	if err != nil {
		_ = tx.Rollback() // 回滚
		fmt.Println("updata id5 err", err)
		return
	}
	sqlstr = "update user set age=age-? where id=?"
	_, err = tx.Exec(sqlstr, 2, 6)
	if err != nil {
		_ = tx.Rollback() // 回滚
		fmt.Println("updata id6 err", err)
		return
	}
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
		fmt.Println("initDB failed:", err)
		return
	}
	// insertRowDemo()
	// deleteRowDemo(2)
	// updateRowDemo(55, 3)
	// queryRowDemo(4)
	// queryMultiRowDemo()
	// prepareQueryDemo()
	// prepareInsertDemo()
	transationDemo()
}
