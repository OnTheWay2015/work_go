package game

import (
	"net/http"
	"test_gin/gin"
	"test_gin/utils"
)

func ResVediosInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		//ptp := c.Param("ptp")
		ptp := c.DefaultQuery("ptp", "")
		res := G_vedios.getVediosInfo(utils.Atoint32(ptp))
		c.JSON(http.StatusOK, gin.H{"res": 1, "data": res})
	}
}
