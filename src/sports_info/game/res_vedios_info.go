package game

import (
	"net/http"
	"sports_info/utils"

	"github.com/gin-gonic/gin"
)

func ResVediosInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		ptp := utils.Atoint32(c.DefaultQuery("ptp", ""))
		spid := utils.Atoint32(c.DefaultQuery("spid", ""))
		res := G_vedios.getVediosInfo(ptp, spid)
		c.JSON(http.StatusOK, gin.H{"res": 1, "data": res, "ptp": ptp, "spid": spid})
	}
}
