package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*")
	engine.Static("/static", "./static")
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	engine.GET("/hello-world", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello World!")
	})
	engine.GET("/json", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"username": "admin",
			"age":      18,
		})
	})
	engine.GET("/xml", func(context *gin.Context) {
		context.XML(http.StatusOK, gin.H{
			"username": "admin",
			"age":      18,
		})
	})
	engine.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"name":  "admin",
			"age":   24,
			"names": []string{"张三", "李四", "王五"},
		})
	})
	engine.GET("/get", func(context *gin.Context) {
		name := context.Query("name")
		age := context.Query("age")
		fmt.Println(name)
		fmt.Println(age)
		context.String(http.StatusOK, "success: name="+name+",age="+age)
	})
	engine.POST("/post", func(context *gin.Context) {
		name := context.PostForm("name")
		age := context.PostForm("age")
		fmt.Println(name)
		fmt.Println(age)
		context.String(http.StatusOK, "success: name="+name+",age="+age)
	})
	engine.GET("/user/:id", func(context *gin.Context) {
		id := context.Param("id")
		context.String(http.StatusOK, "success: id="+id)
	})
	err := engine.Run()
	if err != nil {
		fmt.Println("http server error")
		return
	}
}
