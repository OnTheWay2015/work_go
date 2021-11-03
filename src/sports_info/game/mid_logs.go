package game

import (
	"bytes"
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
)

type BodyLogWriterI interface {
	GetBuf() *bytes.Buffer
}

type BodyLogWriter struct {
	gin.ResponseWriter
	bodyBuf *bytes.Buffer
}

func (w BodyLogWriter) GetBuf() *bytes.Buffer {
	return w.bodyBuf
}
func (w BodyLogWriter) Write(b []byte) (int, error) {
	//memory copy here!
	w.bodyBuf.Write(b)
	return w.ResponseWriter.Write(b)
}

func ReqLogs() gin.HandlerFunc {
	return func(c *gin.Context) {
		buf := bytes.NewBufferString("")
		blw := &BodyLogWriter{c.Writer, buf}
		c.Writer = blw

		//开始时间
		startTime := time.Now()
		//处理请求
		c.Next()
		//结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		//请求方式
		reqMethod := c.Request.Method
		//请求路由
		reqUrl := c.Request.RequestURI
		//状态码
		statusCode := c.Writer.Status()
		//请求ip
		clientIP := c.ClientIP()

		//if statusCode >= 400 {
		//	//ok this is an request with error, let's make a record for it
		//	// now print body (or log in your preferred way)
		//	fmt.Println("Response body: " + str)
		//}
		if G_config.Data.Log.Log_resbody {
			w, res := c.Writer.(BodyLogWriterI)
			bodystr := ""
			if res {
				str := w.GetBuf().String()
				m := map[string]interface{}{}
				json.Unmarshal([]byte(str), &m)

				jj, _ := json.MarshalIndent(m, "", "     ")
				bodystr = string(jj)
			}
			params := c.Params
			paramsStr, _ := json.MarshalIndent(params, "", "     ")
			G_log.Infof("method: %s, uri: %s, status: %d, clientIP: %s, costtm: %d ms, params:%s, resbody:\n%s",
				reqMethod, reqUrl, statusCode, clientIP, latencyTime.Milliseconds(), paramsStr, bodystr)

		} else {

			G_log.Infof("method: %s, uri: %s, status: %d, clientIP: %s, costtm: %d ms ", reqMethod, reqUrl, statusCode, clientIP, latencyTime.Milliseconds())
		}

	}
}
