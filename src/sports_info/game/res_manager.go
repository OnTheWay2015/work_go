package game

import (
	"github.com/gin-gonic/gin"
)

func ResRegs(r *gin.Engine) {
	r.GET("/", ResRoot())

	r.GET("/vedios_info", ResVediosInfo())
	v1 := r.Group("/v1")
	{
		v1.GET("/sports_info/vedios_info", ResVediosInfo())
	}

	//r.LoadHTMLGlob("./*.html") // 指明html加载文件目录
	//r.GET("/apidoc", func(ctx *gin.Context) {
	//	ctx.HTML(http.StatusOK, "apidoc.html", nil)
	//})
}
