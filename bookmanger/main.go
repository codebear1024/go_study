package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"study/bookmanger/models"
)

func bookListHandler(c *gin.Context) {
	// 查数据
	books, err := models.listBook()
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
	err := models.insertBook(title, price)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err.Error(),
		})
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/book/list")
}

func bookDeleteHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	err := models.deleteBook(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err.Error(),
		})
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/book/list")

}

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/sql_test"
	err := models.initDB(dsn)
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
