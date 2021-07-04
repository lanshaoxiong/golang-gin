package main

import (
	"gin"
	"net/http"
)

func main() {
	engine := gin.New()

	// set up pattern and handlerFunc
	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Lan</h1>")
	})

	engine.GET("/hello", func(c *gin.Context) {
		// expect /hello?name=lan
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	engine.GET("/hello/:name", func(c *gin.Context) {
		// expect /hello/lan
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	engine.GET("/assets/*filepath", func(c *gin.Context) {
		// expect /assets/lan/shaoxiong
		c.JSON(http.StatusOK, gin.H{"filepath": c.Param("filepath")})
	})

	engine.POST("/login", func(c *gin.Context) {
		// expect: curl "http://localhost:9999/login" -X POST -d 'username=geektutu&password=1234'
		c.JSON(http.StatusOK, gin.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	// wait for path calling
	engine.Run(":9999")
}
