package symmetric

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// 注：
// aes加密，分组长度128比特，密钥长度以32比特为单位在128比特-256比特之间选择，默认规格中只有128，192，256，
// go提供的接口中，密钥长度只能为16字节
// ecb/ctr:不需要初始化向量（go接口中的iv可以理解为随机数种子, iv的长度 == 明文分组的长度）
// cbc/cfb/ofb:需要初始化向量（des/3des->8字节，aes->16字节，加解密向量相同）
// 如果使用ecb/cbc分组模式需要对明文分组进行填充,cfb/ofb/ctr都不需要对分组明文填充

// 密文先分组，再加密
// des/3des按8字节分组
// 分组密码模式(ecb/cbc模式需要填充 )

// 1、创建des/3des/aes密码接口
// crypto/des   (des/3des)
// crypto/aes   (aes)
// NewCipher(key []byte)(cipher.Block,error)

// 2、如果使用cbc/ecb分组模式需要对密文分组进行填充,cfb/ofb/ctr都不需要对分组密文填充

// 3、创建分组模式接口对象
// cbc:NewCBCEncrypter(b block,iv []byte) BlockMode 加密
// cfb:NewCFBEncrypter(b Block, iv []byte) Stream 加密

func AesEncrypt(orig string, key string) string {
	// 转成字节数组
	origData := []byte(orig)
	k := []byte(key)
	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// block, _ := des.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式 初始化向量默认为一个字节数组的block.BlockSize()长度
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	// 补全码
	origData = PKCS7Padding(origData, blockSize)
	// 创建数组
	cryted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(cryted, origData)
	return base64.StdEncoding.EncodeToString(cryted)
}

func AesDecrypt(cryted string, key string) string {
	// 转成字节数组
	crytedByte, _ := base64.StdEncoding.DecodeString(cryted)
	k := []byte(key)
	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式，初始化向量默认为一个字节数组的block.BlockSize()长度
	//blockMode := cipher.NewCFBDecrypter(block, k[:blockSize])
	//blockMode := cipher.NewOFB(block, k[:blockSize])
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// 创建数组
	orig := make([]byte, len(crytedByte))
	// 解密
	//blockMode.XORKeyStream(orig, crytedByte)
	blockMode.CryptBlocks(orig, crytedByte)
	// 去补全码
	orig = PKCS7UnPadding(orig)
	return string(orig)
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	paddingText := bytes.Repeat([]byte{0}, padding) // 用0填充
	return append(ciphertext, paddingText...)
}

func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimRightFunc(origData, func(r rune) bool {
		return r == rune(0)
	})
}

// PKCS5填充块的大小为8bytes(64位)
// PKCS7填充块的大小可以在1-255bytes之间。
// PKCS7是兼容PKCS5的，PKCS5相当于PKCS7的一个子集

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	paddingText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, paddingText...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}
