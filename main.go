package main

import (
	"gin"
	"net/http"
)

func main() {
	engine := gin.New()
	engine.GET("/", func(c *gin.Context) {
		// fmt.Fprintf(c.Writer, "URL.Path = %q\n", c.Req.URL.Path)
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	engine.GET("/hello", func(c *gin.Context) {
		// for k, v := range c.Req.Header {
		// 	fmt.Fprintf(c.Writer, "Header[%q] = %q\n", k, v)
		// }
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)

	})

	engine.POST("/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	engine.Run(":9999")
}
