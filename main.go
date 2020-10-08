package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello world")
	// 初始化引擎
	r := gin.Default()

	// 注册路由
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 运行，默认开启 8080端口，也可以自定义端口
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
