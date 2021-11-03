package test__123

import "github.com/gin-gonic/gin"

func Test_gin() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello")
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
