package game

import (
	"test_gin/gin"
)

func MidRegs(r *gin.Engine) {
	r.Use(ReqLogs())
}
