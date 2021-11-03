package crypto_aux

import (
	"encoding/base64"
	"strings"
)

func BaseStdEncode(srcBtye []byte) string {
	encoding := base64.StdEncoding.EncodeToString(srcBtye)
	return encoding
}
func BaseDeEncode(src string) string {
	reader := strings.NewReader(src)
	decoder := base64.NewDecoder(base64.StdEncoding, reader)
	// 以流式解码
	buf := make([]byte, 2)
	// 保存解码后的数据
	dst := ""
	for {
		n, err := decoder.Read(buf)
		if n == 0 || err != nil {
			break
		}
		dst += string(buf[:n])
	}
	return dst
}
