package test__123

import (
	"helloworld/controllers"
	_ "helloworld/routers"

	"github.com/astaxie/beego"
)

//=================================================================
//在 conf/app.conf 配置里配置
/*
######################基础配置############################
AppName             = app test
HttpPort            = 8099     #端口
RunMode             = dev
CommentRouterPath ="helloworld\controllers"   #注解路由文件目录

*/

//=================================================================
//https://beego.me/docs/mvc/controller/router.md
func Test_beego() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/testfun", &controllers.Test05Controller{}, "get:GetAct")  //可指定  method 和 对应要处理的方法名
	beego.Router("/testfun", &controllers.Test05Controller{}, "post:GetAct") //可指定  method 和 对应要处理的方法名

	beego.Router("/api01/:all", &controllers.Test01Controller{})
	beego.Router("/api02/:id([0-9]+)", &controllers.Test02Controller{})

	//beego.Include(&controllers.Test03Controller{}) //注解路由

	ns := beego.NewNamespace("/v1",
		beego.NSInclude(&controllers.Test03Controller{}))
	beego.AddNamespace(ns)

	beego.Run()
}

/*
还可以通过参数进行过滤，如果匹配参数就执行
beego.Router("/:id([0-9]+)", &admin.EditController{})
beego.FilterParam("id", func(rw http.ResponseWriter, r *http.Request) {
    dosomething()
})
当然你还可以通过前缀过滤

beego.FilterPrefixPath("/admin", func(rw http.ResponseWriter, r *http.Request) {
    dosomething()
})
*/
