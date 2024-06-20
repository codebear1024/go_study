package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"
)

var log = logrus.New()

func bookListHandler(c *gin.Context) {
	// 查数据
	books, err := listBook()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err.Error(),
		})
		return
	}
	// 返回给浏览器
	c.HTML(http.StatusOK, "list.html", books)
}

func bookNewHandler(c *gin.Context) {
	title := c.PostForm("title")
	price, _ := strconv.Atoi(c.PostForm("price"))
	id, err := insertBook(title, price)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err.Error(),
		})
		return
	}
	log.WithFields(logrus.Fields{
		"id":    id,
		"title": title,
		"price": price,
	}).Debug("书籍添加成功")
	c.Redirect(http.StatusMovedPermanently, "/book/list")
}

func bookDeleteHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	err := deleteBook(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err.Error(),
		})
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/book/list")

}

func initLogrus() (err error) {
	log.Formatter = &logrus.JSONFormatter{} //记录JSON格式日志
	// 指定日志输出
	f, _ := os.Create("./gin.log")
	log.Out = f
	// 告诉GIN框架把日志也记到我们打开的文件中
	// gin.SetMode(gin.ReleaseMode) // 上线的时候要设置为releaseMode，该模式下不会打印gin日志
	// gin.DisableConsoleColor()		// 禁用颜色
	gin.DefaultWriter = log.Out
	// 设置日志级别
	log.Level = logrus.DebugLevel
	return
}

func main() {
	// 日志配置初始化
	err := initLogrus()
	if err != nil {
		panic(err)
	}
	// 数据库初始化
	dsn := "root:@tcp(127.0.0.1:3306)/sql_test"
	err = initDB(dsn)
	if err != nil {
		fmt.Println("InitDB err:", err)
		return
	}
	r := gin.Default()
	r.LoadHTMLGlob("view/*")
	r.GET("/book/list", bookListHandler)
	r.GET("/book/new", func(c *gin.Context) {
		c.HTML(http.StatusOK, "new.html", gin.H{})
	})
	r.POST("/book/new", bookNewHandler)
	r.GET("/book/delete", bookDeleteHandler)
	r.Run()
}
