package test__123

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"strings"
)

const (
	aa int = 1
	bb
	cc
	dd
)

func AesEncryptECB(origData []byte, key []byte) (encrypted []byte) {
	//chr, _ := aes.NewCipher(generateKey(key))
	chr, _ := aes.NewCipher(key)
	length := (len(origData) & aes.BlockSize) / aes.BlockSize
	if length <= 0 {
		length = 1
	}
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, origData)
	pad := byte(len(plain) - len(origData))
	for i := len(origData); i < len(plain); i++ {
		plain[i] = pad //把填充的长度值保存下来,在解密时 trim 使用
	}
	encrypted = make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, chr.BlockSize(); bs < len(origData); bs, be = bs+chr.BlockSize(), be+chr.BlockSize() {
		chr.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted
}

func AesDecryptECB(encrypted []byte, key []byte) (decrypted []byte) {
	//chr, _ := aes.NewCipher(generateKey(key))
	chr, _ := aes.NewCipher(key)
	decrypted = make([]byte, len(encrypted))
	//
	for bs, be := 0, chr.BlockSize(); bs < len(encrypted); bs, be = bs+chr.BlockSize(), bs+chr.BlockSize() {
		chr.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return decrypted[:trim]
}

func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

func test_aes() {
	key := "ce31UU8412453879"
	aa := "aa_10exty321"
	enc := AesEncryptECB([]byte(aa), []byte(key))
	aa_64 := baseStdEncode(enc)
	//aa_64 := ""
	//enc := AesEncryptECB([]byte(aa), []byte(key))
	dec_64 := baseDeEncode(string(aa_64[:]))
	dec := AesDecryptECB([]byte(dec_64), []byte(key))
	//dec := ""

	fmt.Println("aa64:", aa_64, "enc:", string(enc[:]), "dec_64:", string(dec_64), "dec:", string(dec))
}

func baseStdEncode(srcBtye []byte) string {
	encoding := base64.StdEncoding.EncodeToString(srcBtye)
	return encoding
}
func baseDeEncode(src string) string {
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

func Say() {
	test_aes()
	fmt.Println("test.say()")
}
