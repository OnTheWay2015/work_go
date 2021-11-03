package game

import (
	"github.com/gin-gonic/gin"
)

func MidRegs(r *gin.Engine) {
	r.Use(ReqLogs())
}
