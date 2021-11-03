package game

import (
	"net/http"
	"sports_info/utils"

	"github.com/gin-gonic/gin"
)

func ResVediosInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		ptp := c.DefaultQuery("ptp", "")
		res := G_vedios.getVediosInfo(utils.Atoint32(ptp))
		c.JSON(http.StatusOK, gin.H{"res": 1, "data": res})
	}
}
