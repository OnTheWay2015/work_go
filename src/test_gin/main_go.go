package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"

	"test_gin/game"
	"test_gin/gin" //模块以 GOROOT 为根路径
)

type msg struct {
	AA int32
	bb string
}

func test_more() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")

		//url参数
		urlname := c.DefaultQuery("urlname", "default_url_name")

		//截取/
		action = strings.Trim(action, "/")

		//-----------------------
		// 接收数据结构 或 json
		//var jsonobj msg
		//// 将request的body中的数据，自动按照json格式解析到结构体
		//if err := c.ShouldBindJSON(&jsonobj); err == nil {
		//	// 返回错误信息
		//	// gin.H封装了生成json数据的工具

		//	c.JSON(http.StatusBadRequest, gin.H{"res": "msg obj ok"})
		//  fmt.Printf("%v", jsonobj)
		//	return
		//}

		//-----------------------
		//json
		jsonobj1 := make(map[string]interface{}) //注意该结构接受的内容
		c.BindJSON(&jsonobj1)                    // jsonobj1["name"] ...

		fmt.Printf("%v", jsonobj1)
		//-----------------------

		c.String(http.StatusOK, "name:"+name+", action:"+action+", urlname:"+urlname)

	})
	r.POST("/xxxpost", func(c *gin.Context) {
		//todo
	})
	r.PUT("/xxxput")

	//表单参数
	r.POST("/form", func(c *gin.Context) {
		types := c.DefaultPostForm("type", "post")
		username := c.PostForm("username")
		password := c.PostForm("userpassword")
		// c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
		c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
	})

	//监听端口默认为8080
	r.Run()

}

func main_test() {
	// 创建路由, gin 框架中采用的路由库是基于 httprouter 做的
	// 默认使用了2个中间件 Logger(), Recovery()
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		//ok
		//c.String(200, "Hello1245")
		//c.JSON(200, gin.H{"aaa": "123", "n": 11})
		//c.JSON(200, msg{AA: 3, bb: "123"}) //结构体转成json时,只有首字母大写的key才会被导出.
		a := msg{1, "123"}
		c.JSON(http.StatusOK, a)
	})
	r.Run() // 默认时listen and serve on 0.0.0.0:8080
	//r.Run(":8000") // listen and serve on 0.0.0.0:8000

}

func loadConfig() {
	f, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println("load config error: ", err)
	}
	var config interface{}
	err = json.Unmarshal(f, &config)
	if err != nil {
		fmt.Println("Para config failed: ", err)
	}

}

func test_config() {
	loadConfig()
}
func main() {
	fmt.Println("main act go version:")
	fmt.Println(runtime.Version())
	//main_test()
	//test_more()
	//test_config()
	//db.Test()

	game.Start()
}

/*

//https://www.topgoer.com/gin%E6%A1%86%E6%9E%B6/gin%E8%B7%AF%E7%94%B1/%E4%B8%8A%E4%BC%A0%E5%A4%9A%E4%B8%AA%E6%96%87%E4%BB%B6.html
//https://www.topgoer.com/gin%E6%A1%86%E6%9E%B6/gin%E8%B7%AF%E7%94%B1/%E4%B8%8A%E4%BC%A0%E5%A4%9A%E4%B8%AA%E6%96%87%E4%BB%B6.html


上传单个文件
multipart/form-data格式用于文件上传

gin文件上传与原生的net/http方法类似，不同在于gin把原生的request封装到c.Request中

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
    <form action="http://localhost:8080/upload" method="post" enctype="multipart/form-data">
          上传文件:<input type="file" name="file" >
          <input type="submit" value="提交">
    </form>
</body>
</html>
package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    //限制上传最大尺寸
    r.MaxMultipartMemory = 8 << 20
    r.POST("/upload", func(c *gin.Context) {
        file, err := c.FormFile("file")
        if err != nil {
            c.String(500, "上传图片出错")
        }
        // c.JSON(200, gin.H{"message": file.Header.Context})
        c.SaveUploadedFile(file, file.Filename)
        c.String(http.StatusOK, file.Filename)
    })
    r.Run()
}

*/
