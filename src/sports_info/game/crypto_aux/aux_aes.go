package crypto_aux

import (
	"crypto/aes"
	"fmt"
)

func AesEncryptECB(origData []byte, key []byte) (encrypted []byte) {
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
	chr, _ := aes.NewCipher(key)
	decrypted = make([]byte, len(encrypted))
	for bs, be := 0, chr.BlockSize(); bs < len(encrypted); bs, be = bs+chr.BlockSize(), bs+chr.BlockSize() {
		chr.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return decrypted[:trim]
}

func Test_aes() {
	key := "94@!#(*13&32!@)("
	//aa := "123456a"
	aa := "admin"
	enc := AesEncryptECB([]byte(aa), []byte(key))
	aa_64 := BaseStdEncode(enc)
	//aa_64 := ""
	//enc := AesEncryptECB([]byte(aa), []byte(key))
	dec_64 := BaseDeEncode(string(aa_64[:]))
	dec := AesDecryptECB([]byte(dec_64), []byte(key))
	//dec := ""

	fmt.Println("aa64:", aa_64, "enc:", string(enc[:]), "dec_64:", string(dec_64), "dec:", string(dec))
}
