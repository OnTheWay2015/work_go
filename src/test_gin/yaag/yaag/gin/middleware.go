package gin

import (
	"bytes"
	"log"
	"strings"
	"test_gin/yaag/yaag"
	"test_gin/yaag/yaag/models"

	"test_gin/yaag/yaag/middleware"

	"test_gin/gin"
)

type logwritebuf interface {
	GetBuf() *bytes.Buffer
}

func Document() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !yaag.IsOn() {
			return
		}
		apiCall := models.ApiCall{}
		middleware.Before(&apiCall, c.Request)
		c.Next()
		if yaag.IsStatusCodeValid(c.Writer.Status()) {
			apiCall.MethodType = c.Request.Method
			apiCall.CurrentPath = strings.Split(c.Request.RequestURI, "?")[0]
			apiCall.ResponseCode = c.Writer.Status()

			headers := map[string]string{}
			for k, v := range c.Writer.Header() {
				log.Println(k, v)
				headers[k] = strings.Join(v, " ")
			}
			//apiCall.ResponseHeader = headers

			apiCall.RequestHeader = map[string]string{}

			w, res := c.Writer.(logwritebuf)
			if res {
				str := w.GetBuf().String()
				apiCall.ResponseBody = str
			} else {
				apiCall.ResponseBody = ""
			}

			go yaag.GenerateHtml(&apiCall)
		}
	}
}
