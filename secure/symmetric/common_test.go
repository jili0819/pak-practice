package symmetric

import (
	"fmt"
	"testing"
)

var (
	orig = "hello worldfenwakngekwanglewaglawnkgawle325253"
	key  = "123456781234567812345678"
)

func TestAesEncrypt(t *testing.T) {
	fmt.Println("原文：", orig)
	encryptCode := AesEncrypt(orig, key)
	fmt.Println("密文：", encryptCode)
}

func TestAesDecrypt(t *testing.T) {
	fmt.Println("原文：", orig)
	encryptCode := AesEncrypt(orig, key)
	fmt.Println("密文：", encryptCode)
	DecryptCode := AesDecrypt(encryptCode, key)
	fmt.Println("原文：", DecryptCode)
}

func TestRsaEncrypt(t *testing.T) {
	fmt.Println("原文：", orig)
	encryptCode, err := RsaEncrypt(orig)
	if err != nil {
		return
	}
	fmt.Println("密文：", string(encryptCode))
	DecryptCode, err := RsaDecrypt(encryptCode)
	if err != nil {
		return
	}
	fmt.Println("原文：", string(DecryptCode))
}
