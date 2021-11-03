package game

import (
	"net/http"
	"test_gin/gin"
)

type msg struct {
	AA int32
	bb string
}

func ResRoot() gin.HandlerFunc {
	return func(c *gin.Context) {

		//name := c.Param("name")
		//fmt.Print(name)

		//url参数
		//urlname := c.DefaultQuery("urlname", "default_url_name") 
		//fmt.Println(urlname)

		//ok
		//c.String(http.StatusOK, "Hello1245")
		//c.JSON(http.StatusOK, gin.H{"aaa": "123", "n": 11})
		//c.JSON(http.StatusOK, msg{AA: 3, bb: "123"}) //结构体转成json时,只有首字母大写的key才会被导出.
		a := msg{1, "123"}
		c.JSON(http.StatusOK, a)

		//c.JSON(http.StatusOK, gin.H{"aaa": "123", "n": 11})

	}
}
