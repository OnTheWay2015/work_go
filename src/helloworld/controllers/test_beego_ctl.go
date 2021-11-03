package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type Json_ST struct {
	Arg string
}

/*
GetString(key string) string   //获得 url 和 form 参数 (form 在 post 请求里有, 对应 postman 里的 body.form-data)
GetInt(key string) (int64, error)
GetBool(key string) (bool, error)
*/

func (this *MainController) Get() {
	v_str := this.GetString("arg")
	v_int, err := this.GetInt("arg")
	if err != nil {
		this.Ctx.WriteString("hello world v_str:" + v_str + ", v_int: trans err")
		return
	}
	v_int_str := strconv.Itoa(v_int)
	this.Ctx.WriteString("hello world v_str:" + v_str + ", v_int:" + v_int_str)

}

func (this *MainController) Post() {
	var jst Json_ST
	jsondata := this.Ctx.Input.RequestBody
	err := json.Unmarshal(jsondata, &jst)

	v_str := this.GetString("arg")
	v_int, err := this.GetInt("arg")
	if err != nil {
		this.Ctx.WriteString("post hello world v_str:" + v_str + ", v_int: trans err")
		return
	}
	v_int_str := strconv.Itoa(v_int)
	this.Ctx.WriteString("post hello world v_str:" + v_str + ", v_int:" + v_int_str + ", jst.arg:" + jst.Arg)
}

/*
beego.Router("/api/:id([0-9]+)", &controllers.RController{})
自定义正则匹配 //匹配 /api/123 :id= 123

beego.Router("/news/:all", &controllers.RController{})
全匹配方式 //匹配 /news/path/to/123.html :all= path/to/123.html

beego.Router("/user/:username([\w]+)", &controllers.RController{})
正则字符串匹配 //匹配 /user/astaxie :username = astaxie

beego.Router("/download/*.*", &controllers.RController{})
*匹配方式 //匹配 /download/file/api.xml :path= file/api :ext=xml

beego.Router("/download/ceshi/*", &controllers.RController{})
*全匹配方式 //匹配 /download/ceshi/file/api.json :splat=file/api.json
*/

type Test01Controller struct {
	beego.Controller
}

func (this *Test01Controller) Get() {
	this.Ctx.WriteString("/api01/:all hello world")
}

//============================
type Test02Controller struct {
	beego.Controller
}

func (this *Test02Controller) Get() {
	this.Ctx.WriteString("/api02/:id([0-9]+) hello world")
}

//============================
type Test05Controller struct {
	beego.Controller
}

//beego.Router("/testfun", &controllers.Test05Controller{}, "post:GetAct") //可指定  method 和 对应要处理的方法名
func (this *Test05Controller) GetAct() {
	this.Ctx.WriteString("test05 GetAct")
}

//============================

/*
注解路由
*/
type Test03Controller struct {
	beego.Controller
}

//func (this *Test03Controller) URLMapping() {
//	this.Mapping("test", this.actMapRouter)
//}

// @router /test/:all [get]
func (this *Test03Controller) ActMapRouter() {
	this.Ctx.WriteString("actMapRouter!")
}

//============================

/*
方法表达式路由
*/

type Test04Controller struct {
	beego.Controller
}

func (this *Test04Controller) Get() {
	this.Ctx.WriteString("func router")
}
