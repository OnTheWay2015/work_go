package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["helloworld/controllers:Test03Controller"] = append(beego.GlobalControllerRouter["helloworld/controllers:Test03Controller"],
        beego.ControllerComments{
            Method: "ActMapRouter",
            Router: "/test/:all",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
